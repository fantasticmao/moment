package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	images []string
	height int
)

func main() {
	flag.IntVar(&height, "h", 120, "height of the bottom subtitle in unit px")
	flag.Parse()
	images = flag.Args()
	checkArgs()

	target, err := stitch(images, height)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Success! The final image is %s\n", target)
	}
}

func checkArgs() {
	if len(images) <= 0 {
		fatal("Error: images not found\nusage: %s image [image ......]\n", os.Args[0])
	}
	if height <= 0 {
		fatal("Error: subtitle height must be positive\n")
	}
}

func fatal(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}
