package server

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *Server) CacheFiles() {
	var files []FileInfo

	err := filepath.Walk(s.uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, FileInfo{
				Name: info.Name(),
				Size: info.Size(),
				Date: info.ModTime(),
			})
		}
		return nil
	})

	if err == nil {
		s.cache.Set("file_list", files, 0)
	} else {
		fmt.Println("Error while refetching files", err)
	}
}
