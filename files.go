package main

import (
	"os"
	"path/filepath"
	"strings"
)

func openFiles(paths []string) ([]*os.File, error) {
	files := make([]*os.File, 0, len(paths))
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}
	return files, nil
}

func closeFiles(files []*os.File) {
	for _, file := range files {
		_ = file.Close()
	}
}

func getFileName(file *os.File) string {
	name := filepath.Base(file.Name())
	if index := strings.LastIndex(name, "."); index != -1 {
		return name[:index]
	} else {
		return name
	}
}

func getFileExtension(file *os.File) string {
	return filepath.Ext(file.Name())
}
