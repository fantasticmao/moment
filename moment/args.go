package moment

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

type Options struct {
	Paths    []string     `short:"p" long:"path" required:"true" description:"The file path of image that need to be stitched"`
	Height   int          `short:"h" long:"height" default:"120" description:"The height of bottom subtitle which should be in unit px"`
	Out      string       `long:"out" default:"." description:"The output directory which is used to save the joined image"`
	ShowHelp func() error `long:"help" description:"Display this help message"`
}

func (o *Options) checkArgs() (err error) {
	if len(o.Paths) <= 0 {
		return errors.New("image paths was not specified")
	}
	if o.Height < 0 {
		return errors.New("subtitle height can't be negative")
	}
	return nil
}

func (o *Options) openFiles() ([]*os.File, error) {
	files := make([]*os.File, 0, len(o.Paths))
	for _, path := range o.Paths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func (o *Options) closeFiles(files []*os.File) error {
	for _, file := range files {
		err := file.Close()
		if err != nil {
			return err
		}
	}
	return nil
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
		log.Fatal(err)
	}
	return opts, err
}
