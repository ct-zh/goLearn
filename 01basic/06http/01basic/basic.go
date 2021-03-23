package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func callback(w http.ResponseWriter, r *http.Request) {
	str := fmt.Sprintf("remote: %s\n", r.RemoteAddr)
	_, _ = io.WriteString(w, str)
}

// {Method:GET
// URL:/
// Proto:HTTP/1.1
// ProtoMajor:1
// ProtoMinor:1
// Header:map[
// 	Accept:[
// 		text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9 ]
// 	Accept-Encoding:[
// 		gzip, deflate, br ]
// 	Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8]
// 	Connection:[keep-alive]
// 	Sec-Fetch-Dest:[document]
// 	Sec-Fetch-Mode:[navigate]
// 	Sec-Fetch-Site:[none]
// 	Sec-Fetch-User:[?1]
// 	Upgrade-Insecure-Requests:[1]
// 	User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36]
// 	]
// Body:{}
// GetBody:<nil>
// ContentLength:0
// TransferEncoding:[]
// Close:false
// Host:127.0.0.1:8777
// Form:map[]
// PostForm:map[]
// MultipartForm:<nil>
// Trailer:map[]
// RemoteAddr:127.0.0.1:53185
// RequestURI:/ TLS:<nil> Cancel:<nil> Response:<nil> ctx:0xc0000224c0}

func main() {
	http.HandleFunc("/", callback)
	err := http.ListenAndServe(":8777", nil)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
