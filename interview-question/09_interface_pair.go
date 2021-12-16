// https://blog.golang.org/laws-of-reflection
// A variable of interface type stores a pair:
//  the concrete value assigned to the variable,
//  and that value's type descriptor.
package main

import (
	"fmt"
	"reflect"
)

type printer interface {
	Print()
}

type S struct {
	A int
}

func (s *S) Print() {
	fmt.Println(s.A)
}

func main() {
	p := getNilPrinter()
	safePrint(p)
	p = getNilInterface()
	safePrint(p)
	p = getRealPrinter()
	safePrint(p)
}

func safePrint(p printer) {
	if p == nil {
		return
	}
	switch reflect.TypeOf(p).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		if reflect.ValueOf(p).IsNil() {
			return
		}
	}

	p.Print()
}

func getNilPrinter() printer {
	var p *S
	return p
}

func getNilInterface() printer {
	return nil
}

func getRealPrinter() printer {
	return &S{3}
}
