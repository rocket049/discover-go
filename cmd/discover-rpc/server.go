package main

import (
	"net"
	"net/rpc"

	"gitee.com/rocket049/discover"
)

type DiscoverRpc struct {
	server *discover.DiscoverServer
}

type ArgRpc struct {
	Scheme string
	IP     string
	Port   int
	Uri    string
	Name   string
	Title  string
}

func (s *DiscoverRpc) Append(arg ArgRpc, ret *int) error {
	s.server.Append(arg.Scheme, arg.IP, arg.Port, arg.Uri, arg.Name, arg.Title)
	return nil
}

func (s *DiscoverRpc) Remove(arg ArgRpc, ret *int) error {
	s.server.Remove(arg.Scheme, arg.IP, arg.Port, arg.Uri)
	return nil
}

func main() {
	svr := discover.NewServer()
	go svr.Serve(true)
	rpcHandler := &DiscoverRpc{server: svr}
	rpc.Register(rpcHandler)

	l, err := net.Listen("tcp", "127.0.0.1:9660")
	if err != nil {
		panic(err)
	}
	rpc.Accept(l)
}
