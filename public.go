package discover

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

type ServeNode struct {
	Href, Title, Name string
}

func createUrl(scheme, port, uri, ip string) (res string) {
	if len(scheme) > 0 {
		res = fmt.Sprintf("%s://%s", scheme, ip)
	} else {
		res = ip
	}
	n, err := strconv.ParseInt(port, 10, 32)
	if err == nil && n > 0 {
		res = fmt.Sprintf("%s:%s", res, port)
	}
	if len(uri) > 0 {
		if strings.HasPrefix(uri, "/") {
			res = fmt.Sprintf("%s%s", res, uri)
		} else {
			res = fmt.Sprintf("%s/%s", res, uri)
		}
	}
	return
}

func xmlEscape(s string) string {
	buf := bytes.NewBufferString("")
	err := xml.EscapeText(buf, []byte(s))
	chk(err)
	//return strings.ReplaceAll(buf.String(), `"`, "&quot;")
	return buf.String()
}

//<append scheme=\"%s\" port=\"%d\" uri=\"%s\" title=\"%s\" name=\"%s\" />\n\r
type appendData struct {
	XMLName xml.Name `xml:"append"`
	Scheme  string   `xml:"scheme,attr"`
	Port    string   `xml:"port,attr"`
	Uri     string   `xml:"uri,attr"`
	Title   string   `xml:"title,attr"`
	Name    string   `xml:"name,attr"`
}

//<remove scheme=\"%s\" port=\"%d\" uri=\"%s\" />\n\r
type removeData struct {
	XMLName xml.Name `xml:"remove"`
	Scheme  string   `xml:"scheme,attr"`
	Port    string   `xml:"port,attr"`
	Uri     string   `xml:"uri,attr"`
}

//"<serve href=\"%s\" title=\"%s\"  name=\"%s\" />\n\r"
type serveData struct {
	XMLName xml.Name `xml:"serve"`
	Href    string   `xml:"href,attr"`
	Title   string   `xml:"title,attr"`
	Name    string   `xml:"name,attr"`
}
