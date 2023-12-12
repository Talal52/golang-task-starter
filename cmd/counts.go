package cmd

import (
	"fmt"
	"golang/structures"
)

func FileReader(name string, routines int) (int, int, int, int) {
	var Lines, Words, Vowels, Punctuations int

	channel := make(chan structures.Counter)

	chunk := len(name) / routines

	fmt.Printf("File Chunks = %v ", routines)
	fmt.Printf("\n")

	for i := 0; i < routines; i++ {
		start := i * chunk
		end := (i + 1) * chunk
		go Count(name[start:end], channel)
	}

	for i := 0; i < routines; i++ {
		Counts := <-channel
		Lines = Lines + Counts.Lines
		Words = Words + Counts.Words
		Vowels = Vowels + Counts.Vowels
		Punctuations = Punctuations + Counts.Punctuations
	}
	return Lines, Words, Vowels, Punctuations
}

func Count(fileContent string, channel chan structures.Counter) {
	count := structures.Counter{}
	for _, char := range fileContent {
		switch {
		case char == ' ' || char == '\t' || char == '\r' || char == '.' || char == ',' || char == ';' || char == ':' || char == '!' || char == '?' || char == '(' || char == ')' || char == '[' || char == ']' || char == '{' || char == '}':
			count.Words++
		case char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u':
			count.Vowels++
		case char == '.' || char == '!' || char == '?' || char == ',' || char == ':' || char == ';' || char == '(' || char == ')' || char == '[' || char == ']' || char == '{' || char == '}':
			count.Punctuations++
		case char == '\n':
			count.Lines++
		}
	}
	channel <- count
}
