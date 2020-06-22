package main

import (
	"flag"

	"gitee.com/rocket049/discover"
)

func main() {
	var readonly = flag.Bool("readonly", false, "readonly mode")
	flag.Parse()
	if *readonly == false {
		discover.Serve()
	} else {
		server := discover.NewServer()
		server.Append("http", "192.168.1.6", "0", "qq", "serveqq", "title of this")
		server.Serve(true)
	}
	//discover.Serve()
	/* 另一种方式:*/

}
