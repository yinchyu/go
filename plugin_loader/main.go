package main

import (
	"fmt"
	"plugin"
)

func main() {
	open, err := plugin.Open("./plugin.so")
	if err != nil {
		fmt.Println(err)

	}
	lookup, err := open.Lookup("LoaderName")
	if err != nil {
		fmt.Println(err)

	}
	loader, ok := lookup.(CommonLoader)
	loader2, ok2 := lookup.(Loader)
	fmt.Println(ok, loader)
	fmt.Println(ok2, loader2)
	fmt.Println(loader.GetError())
	loader.SetError("23234")
	fmt.Println(loader.GetError())
	sloader, ok := lookup.(Loader)
	fmt.Println(ok, sloader)
	fmt.Println(sloader.GetError())
	sloader.SetError("567657")
	fmt.Println(sloader.GetError())
	fmt.Println(sloader.GetLoaderName())
}
