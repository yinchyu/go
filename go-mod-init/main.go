package main

import (
	"fmt"
	"getRouter"
	"router"
	"runRouter"
	"test"
)

func main() {
	fmt.Println("1")
	test.Test()
	fmt.Println("2")
	router.Router()
	fmt.Println("3")
	runRouter.RunRouter()
	fmt.Println("4")
	getRouter.GetRouter()
	fmt.Println("5")
	fmt.Println("1111")
}
