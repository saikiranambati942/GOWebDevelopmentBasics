package main

import ("fmt"
"net/http"
)

type dog int
func (m dog) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"doggy doggy")
}

type rabbit int

func (r rabbit) ServeHTTP(w http.ResponseWriter,r2 *http.Request){
	fmt.Fprintln(w,"rabbit rabbit")
}

func main(){
	var d dog
	var r rabbit
	mux:=http.NewServeMux()
	mux.Handle("/dog/",d)
	mux.Handle("/rabbit",r)
	http.ListenAndServe(":8080",mux)
}
