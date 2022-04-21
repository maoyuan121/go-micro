package main

import (
	"fmt"

	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/file"
)

func main() {
	// 从一个文件 source 加载配置
	if err := config.Load(file.NewSource(
		file.WithPath("./config.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}

	// 定义 host 类型
	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	var host Host

	// 读取一个 database host
	if err := config.Get("hosts", "database").Scan(&host); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(host.Address, host.Port)
}
