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
	email := "1192867487@qq.com"
	code, err := GetCode(email)
	if err == nil {
		if VerifyEmailFormat(email) {
			ok := SendEmail(email, code)
			fmt.Println(ok)
		}
	} else {
		fmt.Println("发送的太过频繁")
	}

}

func TestSend_Ms(t *testing.T) {
	defer ClostClient()
	mobile := "18898186026"
	code, err := GetCode(mobile)
	if err == nil {
		if VerifyTelephoneFormat(mobile) {
			text := fmt.Sprintf("您的验证码是：%s。请不要把验证码泄露给其他人。", code)
			Send_Ms(text, mobile)
		}
	} else {
		fmt.Println("发送的太过频繁")
	}
}
