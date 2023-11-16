package main

import (
	"fmt"
	"os"
)

func checkword(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func vowels(char byte) bool {
	return (char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U')
}

func checkPunctua(char byte) bool {
	return (char >= '!' && char <= '/') || (char >= ':' && char <= '@') ||
		(char >= '[' && char <= '`') || (char >= '{' && char <= '~')
}

func task(filename string) (int, int, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	defer file.Close()
	row := 1
	word := 0
	vowel := 0
	punctuation := 0
	withinWord := false

	for {
		var char [1]byte
		_, err := file.Read(char[:])
		if err != nil {
			break
		}
		if char[0] == '\n' {
			row++
			withinWord = false
		}
		if checkword(char[0]) {
			if !withinWord {
				word++
				withinWord = true
			}
		}
		if vowels(char[0]) {
			if !withinWord {
				vowel++
				withinWord = true
			}
		} else {
			withinWord = false

			if checkPunctua(char[0]) {
				punctuation++
			}
		}
	}

	return row, word, vowel, punctuation, nil
}

func main() {
	filename := "text.txt"

	row, words, vowel, punctuations, err := task(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	total := words + punctuations

	fmt.Printf("total: %d\n", total)
	fmt.Printf("total words: %d\n", words)
	fmt.Printf("total vowels: %d\n", vowel)
	fmt.Printf("total punctuation: %d\n", punctuations)
	fmt.Printf("total rows: %d\n", row)

}
