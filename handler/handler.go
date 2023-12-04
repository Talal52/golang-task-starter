package handler

import (
	"bufio"
	"fmt"
	"golang/files"
	"golang/structures"
	"golang/utils"
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-analyzer",
	Short: "Analyzes a text file and provides statistics",
}

var fileCmd = &cobra.Command{
	Use:   "files",
	Short: "count the number of lines, words",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		limit := args[1]
		AnalyzeFile(fileName, limit)
	},
}

func AnalyzeFile(fileName, limit string) {
	startingTime := time.Now()
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file:", fileName, limit, err)
		return
	}
	defer file.Close()

	const chunkSize = 4000
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
				result := files.Counts(c)
				resultChan <- result
			}(chunk)
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	finalResult := utils.SummarizeResults(resultChan)

	fmt.Printf("Lines: %d\nWords: %d\nVowels: %d\nPunctuations: %d\n",
		finalResult.LineCount, finalResult.WordsCount, finalResult.VowelsCount, finalResult.PunctuationCount)

	endTime := time.Now()
	totalTime := endTime.Sub(startingTime)
	fmt.Printf("Execution time: %v\n", totalTime)
}

func Init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.PersistentFlags().StringP("any", "t", "56", "rt")
}
