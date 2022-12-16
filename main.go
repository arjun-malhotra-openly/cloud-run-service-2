package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":80", nil)
}

func response(res http.ResponseWriter, req *http.Request) {
	responseService2 := "Response From Server 2"
	io.WriteString(res, responseService2)
}
