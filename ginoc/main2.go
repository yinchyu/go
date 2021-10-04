package main

import (
	"expvar"
	"fmt"
	"net/http"
)

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// func main() {
// 	http.HandleFunc("/", handler)
// 	err:=http.ListenAndServe(":80", nil)
// 	if err!=nil{
// 		fmt.Println(err)
// 	}
// }
