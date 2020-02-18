package app

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取文件名称（去除后缀名）
func getBaseNameWithoutExtension(file *os.File) string {
	name := filepath.Base(file.Name())
	if index := strings.LastIndex(name, "."); index != -1 {
		return name[:index]
	} else {
		return name
	}
}

// 获取文件后缀名
func getExtension(file *os.File) string {
	return filepath.Ext(file.Name())
}
