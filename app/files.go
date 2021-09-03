package app

import (
	"os"
	"path/filepath"
	"strings"
)

// GetBaseNameWithoutExtension 获取文件名称（去除后缀名）
func GetBaseNameWithoutExtension(file *os.File) string {
	name := filepath.Base(file.Name())
	if index := strings.LastIndex(name, "."); index != -1 {
		return name[:index]
	} else {
		return name
	}
}

// GetExtension 获取文件后缀名
func GetExtension(file *os.File) string {
	return filepath.Ext(file.Name())
}
