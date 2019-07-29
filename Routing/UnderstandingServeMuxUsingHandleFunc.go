package main
import(
	"fmt"
	"net/http"
)


func p (w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"parrot parrot")
}



func c (w http.ResponseWriter,r2 *http.Request){
	fmt.Fprintln(w,"crow crow")
}

func main(){

	http.HandleFunc("/parrot/",p)
	http.HandleFunc("/crow",c)
	http.ListenAndServe(":8080",nil)
}
