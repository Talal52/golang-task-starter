package handlers

import (
	"fmt"
	"golang/cmd"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func FileData(c *gin.Context) {
	startTime := time.Now()
	var Routines int
	Routines, err := strconv.Atoi(c.Query("Routines"))
	if err != nil {
		fmt.Println("Invalid Go Routines")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Go Routines"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Failed to parse multipart form:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files := form.File["file"]

	if len(files) == 0 {
		fmt.Println("No file parameter found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	file, err := files[0].Open()
	if err != nil {
		fmt.Println("Failed to open file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	totalLines, totalWords, totalVowels, totalPunctuations := cmd.FileReader(string(fileData), Routines)
	totalTime := time.Since(startTime)
	totalTimeString := totalTime.String()
	c.JSON(http.StatusOK, gin.H{
		"Go Routines":        Routines,
		"Execution Time":     totalTimeString,
		"Total Words ":       totalWords,
		"Total Lines":        totalLines,
		"Total Vowels":       totalVowels,
		"Total Punctuations": totalPunctuations,
	})
}
