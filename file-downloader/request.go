package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func get() {
	url := "http://www.yinchangyu.top:10000/gao.mp4"
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	file, err1 := os.Create("download.mp4")
	defer file.Close()
	defer response.Body.Close()
	if err1 != nil {
		log.Println(err1)
	}
	io.Copy(file, response.Body)

}
