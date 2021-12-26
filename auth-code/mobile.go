package auth_code

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func Send_Ms(text, mobile string) {
	sms_send_url := viper.GetString("mobile.sms_send_url")
	account := viper.GetString("mobile.account")
	password := viper.GetString("mobile.password")
	keys := url.Values{}
	keys.Set("account", account)
	keys.Set("password", password)
	keys.Set("content", text)
	keys.Set("mobile", mobile)
	keys.Set("format", "json")
	keys.Set("Accept", "text/plain")
	fmt.Println(keys.Get("mobile"), keys)
	res, err := http.PostForm(sms_send_url, keys)
	defer res.Body.Close()
	if err != nil {
		log.Println(err)
	}
}
func VerifyTelephoneFormat(telephone string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(telephone)
}
