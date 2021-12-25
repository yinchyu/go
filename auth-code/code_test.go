package auth_code

import (
	"fmt"
	"testing"
	"time"
)

func TestCode(t *testing.T) {
	fmt.Println(Gencode())
}
func TestGetCode(t *testing.T) {
	NewClient()
	for i := 0; i < 10; i++ {
		fmt.Println(GetCode("1234"))
		time.Sleep(time.Second * 1)
	}

}
