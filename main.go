package main

import (
	"bufio"
	"fmt"
	"golang/counter"
	"golang/structures"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-analyzer",
	Short: "Analyzes a text file and provides statistics",
	Args:  cobra.ExactArgs(2),
	Run:   analyzeFile,
}

func analyzeFile(cmd *cobra.Command, args []string) {
	startingTime := time.Now()

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	const chunkSize = 3000
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

	endTime := time.Now()
	totalTime := endTime.Sub(startingTime)
	fmt.Printf("Execution time: %v\n", totalTime)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
