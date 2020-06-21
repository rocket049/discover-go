package main

import "gitee.com/rocket049/discover"

func main() {
	discover.Serve()
	/* 另一种方式:
	server := discover.NewServer()
	server.Serve()
	*/
}
