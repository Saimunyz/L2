package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type args struct {
	addresses []string
	flag      []string
}

func TestWget(t *testing.T) {
	t.Run("without O flag", func(t *testing.T) {
		ceases := []args{
			{
				addresses: []string{"https://losst.ru/komanda-wget-linux"},
			},
			{
				addresses: []string{
					"https://losst.ru/komanda-wget-linux",
					"https://losst.ru/komanda-wget-linux",
					"https://losst.ru/komanda-wget-linux",
				},
			},
		}

		for _, testCases := range ceases {
			command := append([]string{"run", "task.go"}, testCases.addresses...)

			_, err := exec.Command("go", command...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting task.go failed: %v", err)
				clearDirs()
				os.Exit(2)
			}

			os.Mkdir("realOut", 0644)
			os.Chdir("realOut")

			_, err = exec.Command("wget", testCases.addresses...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting wget failed: %v", err)
				clearDirs()
				os.Exit(2)
			}
			os.Chdir("..")

			myOut, err := ioutil.ReadDir("losst.ru")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				clearDirs()
				os.Exit(2)
			}

			realOut, err := ioutil.ReadDir("realOut")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				clearDirs()
				os.Exit(2)
			}

			for i := range realOut {
				real := fmt.Sprintf("%s/%s", "realOut/", realOut[i].Name())
				my := fmt.Sprintf("%s/%s", "losst.ru", myOut[i].Name())
				_, err := compareFiles(real, my)
				if err != nil {
					t.Error(err)
					clearDirs()
					os.Exit(2)
				}
			}
			clearDirs()
		}
	})

	t.Run("with O flag", func(t *testing.T) {
		ceases := []args{
			{
				addresses: []string{"https://losst.ru/komanda-wget-linux"},
				flag:      []string{"-O", "some_file"},
			},
			{
				addresses: []string{
					"https://losst.ru/komanda-wget-linux",
					"https://losst.ru/komanda-wget-linux",
					"https://losst.ru/komanda-wget-linux",
				},
				flag: []string{"-O", "file1.html file2.html file3.html"},
			},
		}

		for _, testCases := range ceases {
			command := append([]string{"run", "task.go"}, testCases.flag...)
			command = append(command, testCases.addresses...)

			_, err := exec.Command("go", command...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting task.go failed: %v", err)
				clearDirs()
				os.Exit(2)
			}

			os.Mkdir("realOut", 0644)
			os.Chdir("realOut")

			command = append(testCases.flag, testCases.addresses...)

			_, err = exec.Command("wget", command...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting wget failed: %v", err)
				clearDirs()
				os.Exit(2)
			}
			os.Chdir("..")

			myOut, err := ioutil.ReadDir("losst.ru")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				clearDirs()
				os.Exit(2)
			}

			realOut, err := ioutil.ReadDir("realOut")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				clearDirs()
				os.Exit(2)
			}

			for i := range realOut {
				real := fmt.Sprintf("%s/%s", "realOut/", realOut[i].Name())
				my := fmt.Sprintf("%s/%s", "losst.ru", myOut[i].Name())
				_, err := compareFiles(real, my)
				if err != nil {
					t.Error(err)
					clearDirs()
					os.Exit(2)
				}
			}
			clearDirs()
		}
	})
}

// clearDirs - delete all dirs for test
func clearDirs() {
	os.RemoveAll("realOut")
	os.RemoveAll("losst.ru")
}

// compareFiles - comparing two files and return true/false and error
func compareFiles(filename1, filename2 string) (bool, error) {
	file1, err := ioutil.ReadFile(filename1)
	if err != nil {
		return false, err
	}

	file2, err := ioutil.ReadFile(filename2)
	if err != nil {
		return false, err
	}

	lines1 := strings.Split(string(file1), "\n")
	lines2 := strings.Split(string(file2), "\n")

	for i := range lines2 {
		if lines1[i] != lines2[i] {
			return false, fmt.Errorf("files: %s!= %s\nShould: %s\nGot: %s\n", filename1, filename2, lines1[i], lines2[i])
		}
	}
	return true, nil
}
