package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/sys/unix"
// )

// // get "/bucket" folder size, its used and available space
// func resource(c *gin.Context) {
// 	info, err := getFolderInfo(uploadDir)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get folder info"})
// 		return
// 	}

// 	response := map[string]interface{}{
// 		"folderSize":      info.Size,
// 		"usedSpace":       info.UsedSpace,
// 		"availableSpace":  info.AvailableSpace,
// 		"filesystemType":  info.FilesystemType,
// 		"mountPoint":      info.MountPoint,
// 		"percentageUsage": info.PercentageUsage,
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// type folderInfo struct {
// 	Size            int64  `json:"folderSize"`
// 	UsedSpace       int64  `json:"usedSpace"`
// 	AvailableSpace  int64  `json:"availableSpace"`
// 	FilesystemType  string `json:"filesystemType"`
// 	MountPoint      string `json:"mountPoint"`
// 	PercentageUsage string `json:"percentageUsage"`
// }

// func getFolderInfo(path string) (*folderInfo, error) {
// 	var stat unix.Statfs_t

// 	err := unix.Statfs(path, &stat)
// 	if err != nil {
// 		return nil, err
// 	}

// 	totalSize := stat.Blocks * uint64(stat.Bsize)
// 	usedSpace := (stat.Blocks - stat.Bfree) * uint64(stat.Bsize)
// 	availableSpace := stat.Bfree * uint64(stat.Bsize)
// 	percentageUsage := fmt.Sprintf("%.2f%%", float64(usedSpace)/float64(totalSize)*100)

// 	return &folderInfo{
// 		Size:            int64(totalSize),
// 		UsedSpace:       int64(usedSpace),
// 		AvailableSpace:  int64(availableSpace),
// 		FilesystemType:  string(stat.Fstypename[:]),
// 		MountPoint:      path,
// 		PercentageUsage: percentageUsage,
// 	}, nil
// }
