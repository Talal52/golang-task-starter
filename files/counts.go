package files

import (
	"golang/structures"
)

func Counts(data string) structures.Summary {
	// fmt.Println("processing data:", data)
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
