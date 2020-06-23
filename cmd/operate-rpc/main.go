package main

import (
	"flag"
	"net"
	"net/rpc"
)

type argRpc struct {
	Scheme string
	IP     string
	Port   int
	Uri    string
	Name   string
	Title  string
}

func main() {
	var d = flag.Bool("d", false, "delete")
	var scheme = flag.String("scheme", "http", "scheme http or https")
	var port = flag.Int("port", 0, "port")
	var uri = flag.String("uri", "", "uri path")
	var title = flag.String("title", "about this", "some about this service")
	var name = flag.String("name", "name", "name of service")
	flag.Parse()

	conn, err := net.Dial("tcp", "127.0.0.1:9660")
	if err != nil {
		panic(err)
	}

	rpcClient := rpc.NewClient(conn)

	if *d {
		arg := argRpc{Scheme: *scheme, IP: "example.org", Port: *port, Uri: *uri, Name: "", Title: ""}
		var ret int
		rpcClient.Call("DiscoverRpc.Remove", arg, &ret)
	} else {
		arg := argRpc{Scheme: *scheme, IP: "example.org", Port: *port, Uri: *uri, Name: *name, Title: *title}
		var ret int
		rpcClient.Call("DiscoverRpc.Append", arg, &ret)
	}
}
