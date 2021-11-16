package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed dir
var filesystem1 embed.FS

func main() {
	fsd, err := fs.Sub(filesystem1, "dir")
	if err != nil {
		fmt.Println(err)
	}
	http.Handle("/", http.FileServer(http.FS(fsd)))
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
