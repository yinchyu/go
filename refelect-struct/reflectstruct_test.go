package refelectstruct

import (
	"fmt"
	"testing"
)

type Uer struct {
	Name string `orm:"name"`
	Age  int
}

func TestPrease(t *testing.T) {
	prease := Prease(&Uer{})
	fmt.Printf("%#v", prease)
}
