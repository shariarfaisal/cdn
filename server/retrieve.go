package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) getFile(c *gin.Context) {
	// Extract the file name from the URL path
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file name"})
		return
	}

	// Open the file from the uploads directory
	file, err := os.Open(s.uploadDir + "/" + fileName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	modTime := time.Now()

	// Serve the file to the client
	c.Writer.Header().Set("Content-Disposition", "inline")
	http.ServeContent(c.Writer, c.Request, fileName, modTime, file)
}
