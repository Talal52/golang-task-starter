package utils

import (
	"golang/structures"
)

func SummarizeResults(resultChan <-chan structures.Summary) structures.Summary {
	finalResult := structures.Summary{}
	for result := range resultChan {
		finalResult.LineCount += result.LineCount
		finalResult.WordsCount += result.WordsCount
		finalResult.VowelsCount += result.VowelsCount
		finalResult.PunctuationCount += result.PunctuationCount
	}
	return finalResult
}
