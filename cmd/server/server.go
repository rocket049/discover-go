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
		//'@?@'表示由客户端从接收到的UDP数据中读取IP地址
		server.Append("http", "@?@", 0, "qq", "serveqq", "title of this")
		server.Serve(true)
	}
	//discover.Serve()
	/* 另一种方式:*/

}
