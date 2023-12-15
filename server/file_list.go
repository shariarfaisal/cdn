package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	Date       time.Time `json:"date"`
	SizeString string    `json:"size_string"`
}

type getFilesRequest struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

// byte to human readable format
func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func (s *Server) getFiles(ctx *gin.Context) {
	var req getFilesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	files, size, err := s.getFileInfoList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Sort files by modification time (you can change the criteria)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Date.After(files[j].Date)
	})

	if files == nil {
		files = []FileInfo{}
	}

	result := paginateFiles(files, req.Limit, req.Page)

	// disk size 65.49 TB
	volumeSize := 65.49 * 1024 * 1024 * 1024 * 1024
	info := gin.H{
		"total_storage": ByteCountIEC(int64(volumeSize)),
		"used_storage":  ByteCountIEC(size),
		"free_storage":  ByteCountIEC(int64(volumeSize) - size),
	}

	pagination := GetPaginationDetails(result, len(files), req.Limit, req.Page, info)

	ctx.JSON(http.StatusOK, pagination)
}

func (s *Server) getFileInfoList() ([]FileInfo, int64, error) {
	var files []FileInfo
	var size int64 = 0

	err := filepath.Walk(s.uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		size += info.Size()

		if !info.IsDir() {
			files = append(files, FileInfo{
				Name:       info.Name(),
				Size:       info.Size(),
				Date:       info.ModTime(),
				SizeString: ByteCountIEC(info.Size()),
			})
		}
		return nil
	})

	if err != nil && !os.IsNotExist(err) {
		return nil, 0, err
	}

	return files, size, nil
}

func paginateFiles(files []FileInfo, limit, page int) []FileInfo {
	start := limit * (page - 1)
	end := start + limit

	if start < 0 {
		start = 0
	} else if start > len(files) {
		start = len(files)
	} else if end < 0 {
		end = limit
	}

	if end > len(files) {
		end = len(files)
	}

	return files[start:end]
}
