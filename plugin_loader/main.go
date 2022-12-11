package main

import (
	"fmt"
	"plugin"
	"time"
)

func getLoader() CommonLoader {
	open, err := plugin.Open("./plugin.so")
	if err != nil {
		fmt.Println(err)
	}
	lookup, err := open.Lookup("LoaderName")
	if err != nil {
		fmt.Println(err)
	}
	loader, ok := lookup.(CommonLoader)
	fmt.Println("cast type is ok", ok)

	return loader
}
func main() {
	for {
		loader := getLoader()
		fmt.Println(loader)
		fmt.Println(loader.GetError())
		time.Sleep(time.Second * 1)
	}

}
