package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

var cb *CircuitBreaker

func init() {
	var st Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}
	// 只有连接成功然后顺利读到数据才算是成功
	return body.([]byte), nil
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

//给出了一个默认的字符串，然后通过默认的字符串位置来初始化随机字符串
const charset = "abcdefghijklmnopqrstuvwxyz"

func RandomString(length int) string {
	b := make([]byte, length)
	charsetLen := len(charset)
	for i := range b {
		b[i] = charset[seededRand.Intn(charsetLen)]
	}
	return string(b)
}
func main() {
	//body, err := Get("http://www.google.com/robots.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(RandomString(10))
}
