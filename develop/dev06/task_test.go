package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type args struct {
	files []string
	flags []string
}

func TestCut(t *testing.T) {
	t.Run("siimple test", func(t *testing.T) {
		cases := []args{
			{
				files: []string{"test"},
				flags: []string{"-f", "2"},
			},
			{
				files: []string{"test"},
				flags: []string{"-f", "1,3"},
			},
			{
				files: []string{"test"},
				flags: []string{"-f", "10"},
			},
			{
				files: []string{"test", "test"},
				flags: []string{"-f", "2,10"},
			},
		}

		for _, testCase := range cases {
			command := append([]string{"run", "task.go"}, testCase.flags...)
			command = append(command, testCase.files...)

			myOut, err := exec.Command("go", command...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting test failed: %v", err)
				os.Exit(2)
			}

			command = append(testCase.flags, testCase.files...)

			realOut, err := exec.Command("cut", command...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting test failed: %v", err)
				os.Exit(2)
			}

			for i := range realOut {
				if realOut[i] != myOut[i] {
					fmt.Println("RealOut\n", string(realOut))
					fmt.Println("MyOut\n", string(myOut))

					t.Errorf("%s!=%s\nShould: %s\nGot: %s\nFlags: %v\n", string(realOut[i]), string(myOut[i]), string(realOut[i]), string(myOut[i]), testCase.flags)
					os.Exit(2)
				}
			}
		}
	})

	t.Run("test with -d and -s flags", func(t *testing.T) {
		cases := []args{
			{
				files: []string{"test"},
				flags: []string{"-f", "2", "-d", " "},
			},
			{
				files: []string{"test"},
				flags: []string{"-f", "1,3", "-d", ",", "-s"},
			},
			{
				files: []string{"test"},
				flags: []string{"-f", "2", "-d", ""},
			},
			{
				files: []string{"test", "test"},
				flags: []string{"-f", "1,3,4", "-d", " ", "-s"},
			},
		}

		for _, testCase := range cases {
			command := append([]string{"run", "task.go"}, testCase.flags...)
			command = append(command, testCase.files...)

			myOut, err := exec.Command("go", command...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting test failed: %v", err)
				os.Exit(2)
			}

			command = append(testCase.flags, testCase.files...)

			realOut, err := exec.Command("cut", command...).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Starting test failed: %v", err)
				os.Exit(2)
			}

			for i := range realOut {
				if realOut[i] != myOut[i] {
					fmt.Println("RealOut\n", string(realOut))
					fmt.Println("MyOut\n", string(myOut))

					t.Errorf("%s!=%s\nShould: %s\nGot: %s\nFlags: %v\n", string(realOut[i]), string(myOut[i]), string(realOut[i]), string(myOut[i]), testCase.flags)
					os.Exit(2)
				}
			}
		}
	})
}
