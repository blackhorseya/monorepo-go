package main

import "github.com/blackhorseya/monorepo-go/adapter/shortenurl/cmd"

// @title Shortenurl API
// @version 0.1.0
// @description This is a sample server for shortenurl.
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @BasePath /api
func main() {
	cmd.Execute()
}
