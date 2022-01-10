package main

import "net/http"

type webErr struct {
	code int
	msg  string
}

func NewWebErr(msg string, code ...int) error {
	err := &webErr{code: http.StatusOK, msg: msg}
	if len(code) > 0 {
		err.code = code[0]
	}
	return err
}

func (w *webErr) Error() string {
	return w.msg
}
