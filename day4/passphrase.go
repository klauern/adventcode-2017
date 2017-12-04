package day4

import (
	"io/ioutil"
	"reflect"
	"strings"
)

type passphrase map[string]int

func mapPassphrase(phrase string) passphrase {
	p := make(passphrase)

	words := strings.Split(phrase, " ")

	for _, v := range words {
		p[v]++
	}

	return p
}

// ValidPassPhrase will check a phrase string and determine if it's valid.
func ValidPassPhrase(phrase string) bool {
	p := mapPassphrase(phrase)
	for _, v := range p {
		if v > 1 {
			return false
		}
	}
	return true
}

// Part1Test will check the input.txt to see how many passwords are invalid.
func Part1Test() int {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return -1
	}

	lines := strings.Split(string(f), "\n")
	validPassPhrases := 0
	for _, v := range lines {
		if ValidPassPhrase(v) {
			validPassPhrases++
		}
	}
	return validPassPhrases
}

type phraseAgram []map[rune]int

func isAnagram(first, second map[rune]int) bool {
	return reflect.DeepEqual(first, second)
}

func makeGram(word string) map[rune]int {
	w := make(map[rune]int)
	for _, v := range word {
		w[rune(v)]++
	}
	return w
}

func makePhraseAgram(phrase string) phraseAgram {
	words := strings.Split(phrase, " ")
	gram := make(phraseAgram, len(words))
	for _, v := range words {
		gram = append(gram, makeGram(v))
	}
	return gram
}

// ValidPassPhraseNoAnagrams does the part 2 check where anagrams must have
// all of the characters needed.
func ValidPassPhraseNoAnagrams(str string) bool {

	gram := makePhraseAgram(str)
	for i := 0; i < len(gram)-1; i++ {
		for j := i + 1; j < len(gram); j++ {
			if isAnagram(gram[i], gram[j]) {
				return false
			}
		}
	}
	return true
}

// Part2Test performs the part test.
func Part2Test() int {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return -1
	}

	lines := strings.Split(string(f), "\n")
	validPassPhrases := 0
	for _, v := range lines {
		if ValidPassPhraseNoAnagrams(v) {
			validPassPhrases++
		}
	}
	return validPassPhrases
}
