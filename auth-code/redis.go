package auth_code

import (
	"fmt"
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

type Myerr string

func (m Myerr) Error() string {
	return "  too many  request"
}

func Gencode() string {
	code := make([]byte, 6)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 6; i++ {
		code[i] = byte(rand.Intn(10) + 48)
	}
	return string(code)
}

func StoreCode(uid string, code string) {
	if _, res := client.Set(uid, code, time.Minute*5).Result(); res != nil {
		log.Println(res)
		panic(res)
	}
}

func GetCode(uid string) (string, error) {
	getc := client.Get(uid)
	if scode, ok := getc.Result(); ok != nil {
		fmt.Println("redis  get key is nil", ok)
		code := Gencode()
		StoreCode(uid, code)
		return code, nil
	} else {
		return scode, Myerr("")
	}
}

func ClostClient() {
	client.Close()
}
