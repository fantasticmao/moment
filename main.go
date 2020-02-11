package main

import (
	"fmt"
	"github.com/FantasticMao/moment/app"
)

func main() {
	opts := app.ParseArgs()
	fmt.Printf("%v\n", opts)
}
