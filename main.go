package main

import (
	"fmt"
	"github.com/fantasticmao/moment/app"
)

func main() {
	// 解析参数
	opts := app.ParseArgs()
	// 拼接图片
	target := app.Stitch(opts)
	fmt.Println("Success! The joined image path is:", target)
}
