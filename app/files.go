package app

import (
	"os"
	"path/filepath"
	"strings"
)

// 拼接图片
func Stitch(height int, out string, images []os.File) (string, error) {
	return "", nil
}

// 获取文件名称（去除后缀名）
func GetBaseNameWithoutExt(file os.File) string {
	name := filepath.Base(file.Name())
	if index := strings.LastIndex(name, "."); index != -1 {
		return name[:index]
	} else {
		return name
	}
}

// 获取文件后缀名
func GetExtension(file os.File) string {
	return filepath.Ext(file.Name())
}
