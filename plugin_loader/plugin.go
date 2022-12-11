package main

import (
	"fmt"
)

type CommonLoader interface {
	GetError() string
	SetError(string)
	SetLoaderName(string)
	GetLoaderName() string
}
type Loader struct {
	LoaderName string
	Err        error
}

func (e *Loader) SetLoaderName(s string) {
	e.LoaderName = s
}

func (e *Loader) GetError() string {
	return "good"
}

func (e *Loader) SetError(error string) {
	fmt.Println("set error info")
}
func (e *Loader) GetLoaderName() string {
	return "ycy name is good"
}

var LoaderName = Loader{
	LoaderName: "ycy_loader_name",
	Err:        nil,
}
