package main

import (
	"golang/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/readFile", handlers.FileData)
	r.Run(":3000")
}
