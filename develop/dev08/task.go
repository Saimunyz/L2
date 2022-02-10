package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shirou/gopsutil/v3/process"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// readLine - reads line from std input
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fmt.Printf("minishell:%s$ ", currDir)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// remove '\n'
	input = input[:len(input)-1]

	return input, nil
}

// netcat - make connection with HOST:PORT via tcp/udp
// and sends all from stdin
func netcat(args []string) string {
	var (
		host string
		port string
	)

	protocol := "tcp"

	if len(args) < 2 {
		return "need to specify HOST and PORT\n"
	}

	if len(args) >= 3 && args[0] == "-u" {
		protocol = "udp"
		host = args[1]
		port = args[2]
	} else {
		host = args[0]
		port = args[1]
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.Dial(protocol, addr)
	if err != nil {
		return fmt.Sprintf("netcat: connection failded: %v\n", err)
	}

	sigs := make(chan os.Signal, 1)
	errors := make(chan error, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// read stdin and send to conn
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				errors <- err
				return
			}

			_, err = conn.Write([]byte(input[:len(input)-1]))
			if err != nil {
				errors <- err
				return
			}
		}
	}()

	select {
	case err := <-errors:
		return fmt.Sprintf("%v\n", err)
	case s := <-sigs:
		return fmt.Sprintf("stopped by signal: %v\n", s)
	}
}

// cd - change current directory to path
func cd(path []string) string {
	if len(path) > 1 {
		return "cd: too many arguments\n"
	}

	if len(path) < 1 {
		home := os.Getenv("HOME")
		os.Chdir(home)
	} else {
		os.Chdir(path[0])
	}

	return ""
}

// ps - returns all live processes
func ps() string {
	processes, err := process.Processes()
	if err != nil {
		return fmt.Sprintf("ps: error: %v\n", err)
	}

	out := strings.Builder{}
	out.WriteString("PID\tCMD\n")

	for _, p := range processes {
		name, _ := p.Name()

		out.WriteString(fmt.Sprintf("%d\t%s\n", p.Pid, name))
	}

	return out.String()
}

// kill - kills process by pid
func kill(args []string) string {
	if len(args) < 1 {
		return "kill: usage: kill [-s sigspec ] pid\n"
	}

	var (
		pidToKill   string
		signalTitle string
		signalFlag  bool
		signal      syscall.Signal
	)

	// parsing -s SIGNALL and pid
	for _, word := range args {
		if signalFlag {
			signalFlag = false
			signalTitle = word
			continue
		}

		if strings.Contains(word, "-s") {
			signalFlag = true
		} else {
			pidToKill = word
		}
	}

	// empty -s SIGNAL
	if signalFlag && len(signalTitle) < 1 {
		return fmt.Sprintf("kill: %s: invalid signal specification\n", signalTitle)
	}

	// default signal
	signal = syscall.SIGTERM

	switch signalTitle {
	case "SIGINT":
		signal = syscall.SIGINT
	case "SIGTERM":
		signal = syscall.SIGTERM
	case "SIGQUIT":
		signal = syscall.SIGQUIT
	case "SIGKILL":
		signal = syscall.SIGKILL
	case "SIGHUP":
		signal = syscall.SIGHUP
	}

	processes, err := process.Processes()
	if err != nil {
		return err.Error()
	}

	for _, p := range processes {
		pid := p.Pid

		if fmt.Sprintf("%d", pid) == pidToKill {
			err := p.SendSignal(signal)
			if err != nil {
				return fmt.Sprintf("kill: err: %v\n", err)
			}
			return fmt.Sprintf("kill: %s was %s\n", pidToKill, signal.String())
		}
	}

	return fmt.Sprintf("kill: there is no signal with pid: %s\n", pidToKill)
}

// applyCommand - apply CLI command
func applyCommand(line string) string {
	if len(line) < 1 {
		return ""
	}

	var args []string

	commands := strings.Fields(line)
	if len(commands) > 1 {
		args = commands[1:]
	}

	var outLine string

	switch commands[0] {
	case "cd":
		outLine = cd(args)
	case "pwd":
		outLine, _ = os.Getwd()
		outLine += "\n"
	case "echo":
		outLine = strings.Join(args, " ")
		outLine += "\n"
	case "kill":
		outLine = kill(args)
	case "ps":
		outLine = ps()
	case "netcat":
		outLine = netcat(args)
	default:
		cmd := exec.Command(commands[0], args...)

		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if isPipe {
			cmd.Stdout = &previousOut
		}

		// changes stdin if pipe
		if previousOut.Len() > 0 {
			cmd.Stdin = &previousOut
		}

		// run command
		err := cmd.Run()
		if err != nil {
			return fmt.Sprintf("%v\n", err)
		}

		outLine = previousOut.String()

	}

	// save the result of the previous command
	if isPipe {
		previousOut.Reset()
		previousOut.WriteString(outLine)
	}

	return outLine
}

var (
	// store cmd output
	previousOut bytes.Buffer
	isPipe      bool
)

func minishell() error {
	for {
		line, err := readLine()
		if err != nil {
			return err
		}

		// quit
		if line == "quit" {
			break
		}

		var outLines string

		// for pipe
		//previousOut = new(strings.Builder)
		cmds := strings.Split(line, "|")
		isPipe = true

		// range over commands
		for _, cmd := range cmds {
			if cmd == cmds[len(cmds)-1] {
				isPipe = false
			}
			outLines = applyCommand(cmd)
		}
		fmt.Print(outLines)
	}

	return nil
}

func main() {
	err := minishell()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
