package app

import (
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

func ParseArgs() (opts Options) {
	parser := flags.NewParser(&opts, flags.PassDoubleDash)
	parser.Usage = "[OPTIONS]"
	opts.ShowHelp = func() error {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
		return nil
	}
	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stdout)
		fmt.Printf("\n")
		panic(err)
	}
	return opts
}
