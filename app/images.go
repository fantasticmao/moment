package app

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

const targetFileNameSuffix = "_final"

// 拼接图片
func Stitch(opts Options) (string, error) {
	// 校验命令行参数
	if err := checkArgs(opts); err != nil {
		return "", err
	}
	// 读取文件
	files := make([]*os.File, 0, len(opts.Paths))
	for _, path := range opts.Paths {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close() // FIXME

		files = append(files, file)
	}

	config, _, err := image.DecodeConfig(files[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height)

	// 最终生成的文件路径
	targetName := getBaseNameWithoutExt(files[0]) + targetFileNameSuffix + getExtension(files[0])
	targetPath, _ := filepath.Abs(opts.Out + "/" + targetName)
	return targetPath, nil
}
