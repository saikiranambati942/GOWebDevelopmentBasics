package main

import (
	"fmt"
	"net/http"
)
type hotdog int
func (m hotdog)ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch r.URL.Path {
	case"/dog":
fmt.Fprintln(w,"doggy doggy doggy")

	case "/cat":
		fmt.Fprintln(w,"kitty kitty kitty")
	}
}
func main(){
	var d hotdog
	http.ListenAndServe(":8080",d)

}
