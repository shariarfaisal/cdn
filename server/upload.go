package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *Server) uploadFile(c *gin.Context) {
	// Parse the form data to retrieve the uploaded file
	err := c.Request.ParseMultipartForm((1024 * 100) << 20) // 100 GB max file size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
		return
	}

	// Retrieve the file from the form data
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		err := fmt.Errorf("file is not present in the form data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	restrictedExt := []string{
		".exe",
		".msi",
		".bat",
		".cmd",
		".com",
		".dll",
		".sys",
		".vbs",
		".vbe",
		".js",
		".jse",
		".jar",
		".scr",
		".pif",
		".ws",
		".wsf",
		".wsc",
		".wsh",
		".ps1",
		".ps1xml",
		".ps2",
		".ps2xml",
		".psc1",
		".psc2",
		".msh",
		".msh1",
		".msh2",
		".mshxml",
		".msh1xml",
		".msh2xml",
		".scf",
		".lnk",
		".inf",
		".reg",
		".doc",
		".xls",
		".ppt",
		".docm",
		".dotm",
		".xlsm",
		".xltm",
		".xlam",
		".html",
		".exe",
	}

	fileExt := filepath.Ext(handler.Filename)

	for _, ext := range restrictedExt {
		if ext == fileExt {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File type not allowed"})
			return
		}
	}

	// Create a new file in the uploads directory
	f, err := os.Create(s.uploadDir + "/" + handler.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create the file"})
		return
	}
	defer f.Close()

	// Copy the uploaded file to the newly created file
	_, err = io.Copy(f, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to copy the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": handler.Filename})
}
