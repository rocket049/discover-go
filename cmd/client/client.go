package main

import (
	"flag"
	"fmt"

	"gitee.com/rocket049/discover"
)

func main() {
	var a = flag.Bool("a", false, "append")
	var d = flag.Bool("d", false, "delete")
	var scheme = flag.String("scheme", "http", "scheme http or https")
	var port = flag.Int("port", 0, "port")
	var uri = flag.String("uri", "", "uri path")
	var title = flag.String("title", "about this", "some about this service")
	var name = flag.String("name", "name", "name of service")
	flag.Parse()
	client := discover.NewClient()
	if client == nil {
		panic("Fail create client")
	}
	if *a {
		client.Append(*scheme, *port, *uri, *name, *title)
		return
	}
	if *d {
		client.Remove(*scheme, *port, *uri)
		return
	}

	services := client.Query()
	for i := range services {
		fmt.Println(services[i].Href, services[i].Name, services[i].Title)
	}
}
