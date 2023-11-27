package counter

import (
	"golang/structures"
)

func Counts(data string) structures.Summary {
	result := structures.Summary{}
	for _, char := range data {
		if char == '.' {
			result.LineCount++
		} else if char == ' ' {
			result.WordsCount++
		} else if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			result.VowelsCount++
		} else if (char >= 33 && char <= 47) || (char >= 58 && char <= 64) || (char >= 91 && char <= 96) || (char >= 123 && char <= 126) {
			result.PunctuationCount++
		}
	}
	return result
}

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
