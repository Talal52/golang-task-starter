// main.go
package main

import (
	"bufio"
	"fmt"
	"golang/counter"
	"golang/structures"
	"os"
	"sync"
	"time"
)

func main() {
	startingtime := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go file.txt")
		return
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	const chunkSize = 1000
	scanner := bufio.NewScanner(file)
	resultChan := make(chan structures.Summary)
	var wg sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i += chunkSize {
			end := i + chunkSize
			if end > len(line) {
				end = len(line)
			}
			chunk := line[i:end]

			wg.Add(1)

			go func(c string) {
				defer wg.Done()
				result := counter.Counts(c)
				resultChan <- result
			}(chunk)
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	finalResult := counter.SummarizeResults(resultChan)

	fmt.Printf("Lines: %d\nWords: %d\nVowels: %d\nPunctuations: %d\n",
		finalResult.LineCount, finalResult.WordsCount, finalResult.VowelsCount, finalResult.PunctuationCount)

	endtime := time.Now()
	totaltime := endtime.Sub(startingtime)
	fmt.Printf("execution time:%v", totaltime)

}
