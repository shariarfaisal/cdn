// main.go
package main

import (
	"cdn/server"
	"os"
)

const uploadDir = "./bucket"

func main() {
	// Create the upload directory if it doesn't exist
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// Create a new Gin router
	s := server.NewServer(":5051")
	s.Start()
}
