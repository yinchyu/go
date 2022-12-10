package main

import "errors"

type CommonLoader interface {
	getError() string
	setError(string)
	setLoaderName(string)
}
type Loader struct {
	LoaderName string
	Err        error
}

func (e Loader) setLoaderName(s string) {
	e.LoaderName = s
}

func (e Loader) getError() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func (e Loader) setError(error string) {
	e.Err = errors.New(error)
}
func (e Loader) getLoaderName() string {
	return "HELLO"
}

var LoaderName = Loader{
	LoaderName: "ycy_loader_name",
	Err:        nil,
}
