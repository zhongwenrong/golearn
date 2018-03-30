package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8888", nil)
}
func handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form["name"]) > 0 {
		fmt.Fprint(rw, "hello world", req.Form["name"][0])
	} else {
		fmt.Fprint(rw, "hello world")
	}
}
