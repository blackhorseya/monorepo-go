package main

import (
	"github.com/blackhorseya/monorepo-go/adapter/redpacket/cmd"
)

// @title RedPacket API
// @version 0.1.0
// @description This is a sample server for redpacket.
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @BasePath /api
func main() {
	cmd.Execute()
}
