package app

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Paths    []string     `short:"p" long:"path" required:"true" description:"The file path of image that need to be stitched"`
	Height   int          `short:"h" long:"height" default:"120" description:"The height of bottom subtitle which should be in unit px"`
	Out      string       `long:"out" default:"." description:"The output directory which is used to save the joined image"`
	ShowHelp func() error `long:"help" description:"Display this help message"`
}

func ParseArgs() (opts Options, err error) {
	parser := flags.NewParser(&opts, flags.PassDoubleDash)
	// see https://github.com/jessevdk/go-flags/issues/240#issuecomment-321556869
	parser.Usage = "[OPTIONS]"
	opts.ShowHelp = func() error {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
		return nil
	}
	// 解析命令行参数
	if _, err = parser.Parse(); err != nil {
		parser.WriteHelp(os.Stdout)
		fmt.Printf("\n")
		panic(err)
	}
	// 校验命令行参数
	err = checkArgs(opts)
	return opts, err
}

func checkArgs(opts Options) (err error) {
	if len(opts.Paths) <= 0 {
		return errors.New("'--path' argument is required")
	}
	if opts.Height < 0 {
		return errors.New("height could not be negative")
	}
	return nil
}
