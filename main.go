package main

import (
	"fmt"
	"golang/handler"
	"os"
)

func main() {
	handler.Init()
	if len(os.Args) < 4 {
		fmt.Println("go run main.go files <filepath> <no.of goroutines>")
		os.Exit(1)
	}
	fileName := os.Args[2]
	limit := os.Args[3]
	handler.AnalyzeFile(fileName, limit)
}
