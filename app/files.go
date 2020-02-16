package app

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const targetFileNameSuffix = "_final"

// 拼接图片
func Stitch(height int, out string, files []os.File) (string, error) {
	// 最终生成的文件路径
	targetFileName := GetBaseNameWithoutExt(files[0]) + targetFileNameSuffix + GetExtension(files[0])
	targetFilePath, _ := filepath.Abs(out + "/" + targetFileName)
	fmt.Println(targetFilePath)
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
	ext := filepath.Ext(file.Name())
	return strings.ToLower(ext)
}

// 获取文件魔数
func GetMagicNumber(file os.File) string {
	bytes := make([]byte, 1<<2)
	if _, err := file.Read(bytes); err == nil {
		magicNumber := hex.EncodeToString(bytes)
		return strings.ToUpper(magicNumber)
	}
	return ""
}
