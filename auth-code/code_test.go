package auth_code

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	NewClient()
}

func TestCode(t *testing.T) {
	ttl := client.TTL("1234")
	result, err := ttl.Result()
	if err != nil {
		return
	}
	fmt.Println(result)

	fmt.Println(Gencode())
}
func TestGetCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetCode("1234"))
		time.Sleep(time.Second * 1)
	}

}

func TestSendEmail(t *testing.T) {
	defer ClostClient()
	code, err := GetCode("1234")
	email := "1192867487@qq.com"
	if err == nil {
		if VerifyEmailFormat(email) {
			ok := SendEmail(email, code)
			fmt.Println(ok)
		}
	} else {
		fmt.Println("发送的太过频繁")
	}

}
