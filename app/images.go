package app

import (
	"os"
	"strings"
)

var imageMagicNumber = map[string]string{
	".bmp":  "424D",
	".jpg":  "FFD8FFDB",
	".jpeg": "FFD8FF",
	".png":  "89504E47",
}

// 校验图片是否合法
func CheckIllegalImage(file os.File) bool {
	fileExt := GetExtension(file)
	fileMagicNumber := GetMagicNumber(file)
	if magicNumber, ok := imageMagicNumber[fileExt]; ok {
		return strings.HasPrefix(fileMagicNumber, magicNumber)
	} else {
		return false
	}
}
