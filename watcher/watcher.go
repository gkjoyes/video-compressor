package watcher

import (
	"log"
	"os"
	"path/filepath"
)

// Visit travels through each sub folder.
func Visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		// error check.
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// filer mp4 files.
		if filepath.Ext(path) == ".mp4" {
			*files = append(*files, path)
		}
		return nil
	}
}
