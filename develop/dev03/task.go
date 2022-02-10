package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Args - all keys params for sort func
type Args struct {
	k int
	n bool
	r bool
	u bool
	M bool

	files []string
}

// splitByWords - splits each line on words
func splitByWords(w []string) [][]string {
	words := make([][]string, len(w))

	for i := range w {
		words[i] = strings.Fields(w[i])
	}

	return words
}

// applyReverseSort - reversing slice
func applyReverse(lines []string) {
	left, right := 0, len(lines)-1

	for left <= right {
		lines[left], lines[right] = lines[right], lines[left]
		left++
		right--
	}
}

// applyStringSort - sorting strings by given column
func applyStringSort(lines []string, column int) {
	var index int

	// sort.Strings(lines)

	if column > 0 {
		index = column - 1

		forSorting := moveUp(lines, func(str []string) bool {
			return len(str) > index
		})

		sort.Slice(forSorting, func(i, j int) bool {
			words := splitByWords(forSorting)
			return words[i][index] < words[j][index]
		})
	}
}

// ftAtoi - converts string digits to int
func ftAtoi(w string) int {
	var (
		i, sign int
		res     int64
	)

	sign = 1

	// skipping spaces
	for i < len(w) && w[i] == ' ' {
		i++
	}

	//skipping signs
	if i < len(w) && (w[i] == '-' || w[i] == '+') {
		if w[i] == '-' {
			sign = -1
		}
		i++
	}

	// converting from string to int
	for i < len(w) && w[i] >= '0' && w[i] <= '9' {
		res = res*int64(10) + int64(w[i]-'0')
		i++
	}

	return int(res * int64(sign))
}

/*
	moveUp - moves up all lines that return true from th compare function
	and returns a subslice starting from the unmoved lines
*/
func moveUp(lines []string, compare func(str []string) bool) []string {
	words := splitByWords(lines)
	min := 0

	for i := range words {
		if !compare(words[i]) {
			lines[min], lines[i] = lines[i], lines[min]
			min++
		}
	}
	sort.Strings(lines[:min])
	return lines[min:]
}

// applyDigitSort - sorting by digits
func applyDigitSort(lines []string, column int) {
	var index int

	if column > 0 {
		index = column - 1
	}

	forSorting := moveUp(lines, func(str []string) bool {
		return len(str) > index && ftAtoi(str[index]) > 0
	})

	sort.Slice(forSorting, func(i, j int) bool {
		words := splitByWords(forSorting)
		return ftAtoi(words[i][index]) < ftAtoi(words[j][index])
	})
}

// applyUnique - remove duplicates
func applyUnique(lines []string) []string {
	unique := make(map[string]bool)

	for _, line := range lines {
		unique[line] = true
	}

	result := make([]string, 0, len(unique))
	for line := range unique {
		result = append(result, line)
	}

	return result
}

// parseDate - parsing date from string
func parseDate(word string) (time.Time, error) {
	t, err := time.Parse("Jan", word)
	if err == nil {
		return t, nil
	}

	t, err = time.Parse("January", word)
	if err == nil {
		return t, nil
	}
	return time.Time{}, err
}

// applyMonthSort - sorting by months
func applyMonthSort(lines []string, column int) {
	var index int

	if column > 0 {
		index = column - 1
	}

	forSorting := moveUp(lines, func(str []string) bool {
		if len(str) > index {
			_, err := parseDate(str[index])
			if err == nil {
				return true
			}
		}
		return false
	})

	sort.Slice(forSorting, func(i, j int) bool {
		words := splitByWords(forSorting)
		t1, _ := parseDate(words[i][index])
		t2, _ := parseDate(words[j][index])
		return t1.Before(t2)
	})
}

// sortLines - sorting lines by specified args
func sortLines(args *Args, words []string) []string {
	if args.u {
		words = applyUnique(words)
	}

	sort.Strings(words)

	if args.n {
		applyDigitSort(words, args.k)
	} else if args.M {
		applyMonthSort(words, args.k)
	} else {
		applyStringSort(words, args.k)
	}

	if args.r {
		applyReverse(words)
	}

	return words
}

// getLines - returns read lines from files/stdin
func getLines(args *Args) ([]string, error) {
	var lines []string

	if len(args.files) > 0 {
		// reading from files
		for _, val := range args.files {
			file, err := os.Open(val)
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
		}
	} else {
		// reading from stdin
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}

			// remove '\n'
			line = line[:len(line)-1]
			lines = append(lines, line)
		}
	}

	return lines, nil
}

// getFlasgs - return all parsed flags
func getFlags() (*Args, error) {
	k := flag.Int("k", 0, "define on witch column apply sort")
	n := flag.Bool("n", false, "activate sorts on digits")
	u := flag.Bool("u", false, "outputs only unique strings")
	r := flag.Bool("r", false, "reverse sorting")
	M := flag.Bool("M", false, "compare (unknown) < `JAN' < ... < `DEC'")

	flag.Parse()

	args := &Args{
		k: *k,
		n: *n,
		u: *u,
		r: *r,
		M: *M,
	}

	if args.k < 0 {
		return nil, errors.New("counter can't be negative")
	}

	args.files = append(args.files, flag.Args()...)

	return args, nil
}

// Sort - sorts words from files
func Sort() (string, error) {
	args, err := getFlags()
	if err != nil {
		return "", err
	}

	lines, err := getLines(args)
	if err != nil {
		return "", err
	}

	sorted := sortLines(args, lines)

	return strings.Join(sorted, "\n"), nil
}

func main() {
	sorted, err := Sort()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(sorted)
}
