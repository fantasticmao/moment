package main

import (
	"fmt"
	"github.com/FantasticMao/moment/moment"
)

func main() {
	// 解析参数
	opts, err := moment.ParseArgs()
	if err != nil {
		panic(err)
	}

	// 拼接图片
	target, err := moment.Stitch(opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success! The joined image path is:", target)
}
