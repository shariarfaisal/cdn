package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) deleteFile(c *gin.Context) {
	// Extract the file name from the URL path
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file name"})
		return
	}

	// delete the file from the uploads directory

	err := os.Remove(s.uploadDir + "/" + fileName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
