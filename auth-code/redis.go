package auth_code

import (
	"github.com/go-redis/redis"
	"log"
	"math/rand"
	"time"
)

var client *redis.Client

func NewClient() {
	ops := &redis.Options{
		Addr:     "42.193.190.143:6388",
		Password: "",
		DB:       0,
	}
	client = redis.NewClient(ops)
}

func Gencode() string {
	code := make([]byte, 6, 6)
	for i := 0; i < 6; i++ {
		code[i] = byte(rand.Intn(10) + 48)
	}
	return string(code)
}

func StoreCode(uid string, code string) {
	if _, res := client.Set(uid, code, time.Second*5).Result(); res != nil {
		log.Println(res)
		panic(res)
	}
}

func GetCode(uid string) string {
	getc := client.Get(uid)
	if scode, ok := getc.Result(); ok != nil {
		code := Gencode()
		StoreCode(uid, code)
		return code
	} else {
		return scode
	}
}
