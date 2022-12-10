package main

import (
	"fmt"
	"plugin"
)

func main() {
	open, err := plugin.Open("main.so")
	if err != nil {
		return
	}
	lookup, err := open.Lookup("LoaderName")
	if err != nil {
		return
	}
	loader, ok := lookup.(CommonLoader)
	fmt.Println(ok, loader)

	fmt.Println(loader.getError())
	loader.setError("23234")
	fmt.Println(loader.getError())
	sloader, ok := lookup.(Loader)
	fmt.Println(ok, sloader)
	fmt.Println(sloader.getError())
	sloader.setError("567657")
	fmt.Println(sloader.getError())
	fmt.Println(sloader.getLoaderName())
}
