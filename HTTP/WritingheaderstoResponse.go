package main

import (
	"fmt"
	"net/http"
)

type hot int

func (m hot) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("sai-key", "This is from sai")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w,"<h1>Any code you write in this func</h1>")
}
func main() {
var d hot
http.ListenAndServe(":8080",d)
}
