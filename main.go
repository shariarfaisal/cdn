// main.go
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

const uploadDir = "./bucket"

func main() {
	// Create the upload directory if it doesn't exist
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// Create a new Gin router
	r := gin.Default()

	// Set up routes
	r.POST("/", handleUpload)
	// r.GET("/resource", resource)
	r.GET("/:filename", handleRetrieve)
	r.DELETE("/:filename", deleteFile)

	// Start the Gin server
	port := ":5051"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := r.Run(port); err != nil {
		panic(err)
	}
}
