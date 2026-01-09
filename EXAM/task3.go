package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
    if len(word) <= 1 {
        return word
    }

    runes := []rune(word)
    first := runes[0]
    
    restRunes := runes[1:]
    for i, j := 0, len(restRunes)-1; i < j; i, j = i+1, j-1 {
        restRunes[i], restRunes[j] = restRunes[j], restRunes[i]
    }
    
    return string(first) + string(restRunes)
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)
	encryptedWords := make([]string, len(words))

	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	testPhrases := []string{
		"Pepe Schnele is a legend",
		"Salam aleykum",
		"a",
		"one two three",
		"один два три",
	}

	for _, phrase := range testPhrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("input: %s\noutput: %s\n\n", phrase, encrypted)
	}
}
