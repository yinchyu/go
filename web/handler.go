package web

import (
	"fmt"
	"log"
	"net/http"
)

func Run(){
	http.HandleFunc("/",HeaderIndex)
	http.HandleFunc("/count",HeaderRange)
	log.Fatal(http.ListenAndServe(":8080",nil))
	// serve 和ListenAndServe  两个 一个是需要自己定义一个tcp listener  一个是直接传递一个端口就可以了
}
func HeaderIndex(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"the url path is %s ",req.URL.Path)
}
func HeaderRange(w http.ResponseWriter, req *http.Request){
	for k,v:=range req.Header{
		// %q 是什么意思
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

}
