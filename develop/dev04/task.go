package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// filter - remove short sets and repeated words
func filter(tmp map[string][]string) {
	unique := make(map[string]bool)

	for key, value := range tmp {
		// remove short sets
		if len(value) < 2 {
			delete(tmp, key)
		}

		// remove repeated words
		for i := range value {
			if !unique[value[i]] {
				unique[value[i]] = true
			} else {
				value[i] = value[len(value)-1]
				tmp[key] = value[:len(value)-1]
			}
		}

	}

}

// getLettersInAlphabetOrder - returns sorted in alphabetorder string
func getLettersInAlphabetOrder(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)

	return strings.Join(letters, "")
}

// makeTemporaryMapOfAnagrms - makes temporary map of anagrams
func makeTemporaryMapOfAnagrms(dictionary []string) map[string][]string {
	tmp := make(map[string][]string)

	// filling map in words
	for _, val := range dictionary {
		loweredWord := strings.ToLower(val)

		letters := getLettersInAlphabetOrder(loweredWord)

		tmp[letters] = append(tmp[letters], loweredWord)
	}

	return tmp
}

func findAnagrams(dictionary []string) map[string][]string {
	tmp := makeTemporaryMapOfAnagrms(dictionary)

	filter(tmp)

	anagrams := make(map[string][]string, len(tmp))

	for _, value := range tmp {
		sort.Strings(value)
		anagrams[value[0]] = value
	}

	return anagrams
}

func main() {

	dictionary := []string{
		"Пятак",
		"Пятак",
		"пятка",
		"Тяпка",
		"слиток",
		"слиток",
		"столик",
		"листок",
		"Топот",
		"Потоп",
	}

	anagrams := findAnagrams(dictionary)

	for key, value := range anagrams {
		fmt.Printf("Key: %s\nValue: %v\n\n", key, value)
	}
}
