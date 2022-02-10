package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Args - stores all arguments
type Args struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool

	pattern string
	files   []string
}

// getArgs - returns *Args struct with parsed flags
func getArgs() (*Args, error) {
	A := flag.Int("A", 0, "Print NUM lines of trailing context after matching lines")
	B := flag.Int("B", 0, "Print NUM lines of leading context before matching lines")
	C := flag.Int("C", 0, "Print NUM lines of output context")
	c := flag.Bool("c", false, "Suppress normal output; instead print a count of matching lines for each input file")
	i := flag.Bool("i", false, "Ignore case distinctions in both the PATTERN and the input files")
	v := flag.Bool("v", false, "Invert the sense of matching, to select non-matching lines")
	F := flag.Bool("F", false, "Interpret PATTERN as a list of fixed strings")
	n := flag.Bool("n", false, "Prefix each line of output with the line number within its input file")

	flag.Parse()

	args := &Args{
		A: *A,
		B: *B,
		C: *C,
		c: *c,
		i: *i,
		v: *v,
		F: *F,
		n: *n,
	}

	if len(flag.Args()) < 1 {
		return nil, errors.New("you need specified 1 argument: pattern")
	}

	pattern := flag.Args()[0]

	if args.i {
		args.pattern = strings.ToLower(pattern)
	} else {
		args.pattern = pattern
	}

	args.files = append(args.files, flag.Args()[1:]...)

	return args, nil
}

// readFile - reads file line by line
func readFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// getMatchedIndexes - returns matched indexes
func getMatchedIndexes(args *Args, lines []string) map[int]bool {
	indexes := make(map[int]bool)

	for i, val := range lines {

		// ignore case
		if args.i {
			val = strings.ToLower(val)
		}

		// treats pattern as string or regex
		if args.F {
			if strings.Contains(val, args.pattern) {
				indexes[i] = true
			}
		} else {
			matched, err := regexp.MatchString(args.pattern, val)
			if err != nil {
				continue
			}
			if matched {
				indexes[i] = true
			}
		}
	}

	return indexes
}

// addNumberOfLine - adding number of lines
func addNumberOfLine(lines []string) []string {
	resultLines := make([]string, len(lines))

	for i := range lines {
		numberedLine := strings.Builder{}

		number := fmt.Sprintf("%d", i+1)
		numberedLine.Grow(len(number) + 1 + len(lines[i]))

		numberedLine.WriteString(number)
		numberedLine.WriteString(":")
		numberedLine.WriteString(lines[i])

		resultLines[i] = numberedLine.String()
	}

	return resultLines
}

// invertIndexes - returns indexes of lines that are not in indexes
func invertIndexes(lines []string, indexes map[int]bool) map[int]bool {
	newIndexes := make(map[int]bool, len(lines)-len(indexes))

	for i := range lines {
		if !indexes[i] {
			newIndexes[i] = true
		}
	}

	return newIndexes
}

// indexToLines - returns lines by indexes
func indexToLines(lines []string, indexes map[int]bool, args *Args) []string {
	var before int
	resultedLines := make([]string, 0, len(indexes))

	for i := range lines {
		if indexes[i] {
			// places a line containing -- between contiguous groups of matches
			if args.B > 0 || args.C > 0 || args.A > 0 {
				if before > 0 && i-before > 1 {
					resultedLines = append(resultedLines, "--")
				}
				before = i
			}

			resultedLines = append(resultedLines, lines[i])
		}
	}

	return resultedLines
}

// addAfterLines - adding amount of lines indexes after each index
func addAfterLines(lines []string, indexes map[int]bool, args *Args) map[int]bool {
	amount := args.A
	newIndexes := make(map[int]bool, len(indexes))

	for i := range indexes {
		newIndexes[i] = true
	}

	for key := range indexes {
		for i := key; i < key+amount+1 && i < len(lines); i++ {
			if !indexes[i] {
				newIndexes[i] = true

				// changing column to minus
				if args.n {
					lines[i] = strings.Replace(lines[i], ":", "-", 1)
				}
			}
		}
	}

	return newIndexes
}

// addBeforeLines - adding amount of lines indexes before each index
func addBeforeLines(lines []string, indexes map[int]bool, args *Args) map[int]bool {
	amount := args.B
	newIndexes := make(map[int]bool, len(indexes))

	for i := range indexes {
		newIndexes[i] = true
	}

	for key := range indexes {
		for i := key; i > key-amount-1 && i >= 0; i-- {
			if !indexes[i] {
				newIndexes[i] = true

				// changing column to minus
				if args.n {
					lines[i] = strings.Replace(lines[i], ":", "-", 1)
				}
			}
		}
	}

	return newIndexes
}

// getLinesByIndexes - returns corresponding lines by index and args
func getLinesByIndexes(args *Args, lines []string, indexes map[int]bool) []string {
	var resultLines []string

	// prints the number of the line from file
	if args.n {
		lines = addNumberOfLine(lines)
	}

	// invert - prints lines that didn't match
	if args.v {
		indexes = invertIndexes(lines, indexes)
	}

	// amount of matched lines
	if args.c {
		return []string{fmt.Sprintf("%d", len(indexes))}
	}

	// prints N lines after and before matched line
	if args.C > 0 {
		args.A = args.C
		args.B = args.C
	}

	// prints N lines after matched line
	if args.A > 0 {
		indexes = addAfterLines(lines, indexes, args)
	}

	// prints N lines before matched line
	if args.B > 0 {
		indexes = addBeforeLines(lines, indexes, args)
	}

	resultLines = indexToLines(lines, indexes, args)

	return resultLines
}

// findLines - return all lines that matches with pattern
func findLines(args *Args) ([]string, error) {
	var resultLines []string

	for _, val := range args.files {

		lines, err := readFile(val)
		if err != nil {
			return nil, err
		}

		matchedIndexes := getMatchedIndexes(args, lines)
		rows := getLinesByIndexes(args, lines, matchedIndexes)

		// adding filename at the beginning of line
		if len(args.files) > 1 {
			for i := range rows {
				namedLines := strings.Builder{}
				namedLines.Grow(len(rows[i]) + len(val) + 1)

				namedLines.WriteString(val)
				namedLines.WriteString(":")
				namedLines.WriteString(rows[i])
				rows[i] = namedLines.String()
			}
		}
		resultLines = append(resultLines, rows...)
	}

	return resultLines, nil
}

// findLinesFromStdin - grep lines from stdin
func findLinesFromStdin(args *Args) ([]string, error) {
	var (
		rows  []string
		lines []string
	)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lines = append(lines, line[:len(line)-1])
	}

	matchedIndexes := getMatchedIndexes(args, lines)
	rows = getLinesByIndexes(args, lines, matchedIndexes)

	return rows, nil
}

// grep - works like linux grep with flags:
// -A -B -C -c -i -v -F -n. For more info man grep
func grep() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("you need specified 1 argument: pattern")
	}

	var (
		lines []string
		err   error
	)

	args, err := getArgs()
	if err != nil {
		return nil, err
	}

	if len(args.files) < 1 {
		lines, err = findLinesFromStdin(args)
	} else {
		lines, err = findLines(args)
	}
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {

	lines, err := grep()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := range lines {
		fmt.Println(lines[i])
	}
}
