package main

import (
	"flag"

	"gitee.com/rocket049/discover"
)

func main() {
	var readonly = flag.Bool("r", false, "readonly mode")
	flag.Parse()
	if *readonly == false {
		discover.Serve()
	} else {
		server := discover.NewServer()
		server.Append("http", "@?@", 6060, "qq", "serveqq", "title of this")
		server.Append("", "@?@", 8001, "", "simple tcp pcm S16LE 20K1CH", "simple-tcp")
		server.Serve(true)
	}
	//discover.Serve()
	/* 另一种方式:*/

}
