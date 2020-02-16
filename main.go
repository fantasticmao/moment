package main

import (
	"fmt"
	"github.com/FantasticMao/moment/app"
	"os"
)

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

	if targetFilePath, err := app.Stitch(opts.Height, opts.Out, files); err == nil {
		fmt.Printf("%v\n", targetFilePath)
	} else {
		panic(err)
	}
}

func closeFile(files []os.File) {
	for _, file := range files {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}
}
