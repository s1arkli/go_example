package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const text = "https://api.github.com/"

func main() {
	resp, err := http.Get(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("%+v\n", resp)
	dump, _ := httputil.DumpResponse(resp, false)
	fmt.Println(string(dump))
	/*
		HTTP/2.0 200 OK
		Content-Length: 297
		Accept-Ranges: bytes
		Content-Type: application/octet-stream
		Date: Fri, 15 May 2026 06:59:07 GMT
		Etag: "69b05381-129"
		Last-Modified: Tue, 10 Mar 2026 17:23:13 GMT
		Server: nginx/1.22.1
		Strict-Transport-Security: max-age=31536000
		X-Tuna-Mirror-Id: nanomirrors
	*/
}
