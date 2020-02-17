package main

import (
	"fmt"
	"github.com/FantasticMao/moment/moment"
)

func main() {
	// 解析参数
	opts := moment.ParseArgs()
	// 拼接图片
	target := moment.Stitch(opts)
	fmt.Println("Success! The joined image path is:", target)
}
