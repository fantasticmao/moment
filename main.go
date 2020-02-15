package main

import (
	"fmt"
	"github.com/FantasticMao/moment/app"
	"os"
	"path/filepath"
)

const targetFileNameSuffix = "_final"

func main() {
	opts, err := app.ParseArgs()
	if err != nil {
		panic(err)
	}

	if len(opts.Paths) == 0 {
		os.Exit(0)
	}

	files := make([]os.File, 0, len(opts.Paths))
	for _, path := range opts.Paths {
		if file, err := os.Open(path); err == nil {
			files = append(files, *file)
		} else {
			panic(err)
		}
	}
	defer closeFile(files)

	fmt.Printf("%v\n", files)

	// 最终生成的文件路径
	targetFileName := app.GetBaseNameWithoutExt(files[0]) + targetFileNameSuffix + app.GetExtension(files[0])
	targetFilePath, _ := filepath.Abs(opts.Out + "/" + targetFileName)
	fmt.Println(targetFilePath)
}

func closeFile(files []os.File) {
	for _, file := range files {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}
}
