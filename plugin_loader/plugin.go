package main

import (
	"errors"
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
	fmt.Println("--", e)
	if e.Err == nil {
		fmt.Println("error is nil--------------")
		return ""
	}
	fmt.Println("check----", e.Err)
	return e.Err.Error()
}

func (e *Loader) SetError(error string) {
	e.Err = errors.New("GDFSGSDGFDSFGSDHFSDHDFHGSG")
	fmt.Println("set error info-----------------")
}
func (e *Loader) GetLoaderName() string {
	return "ycy name is good"
}

var LoaderName = Loader{
	LoaderName: "ycy_loader_name",
	Err:        nil,
}
