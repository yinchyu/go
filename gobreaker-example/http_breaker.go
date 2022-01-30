package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		body, err := ioutil.ReadAll(resp.Body)
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
func add[T ~int | ~float64](a, b T) T {
	return a + b
}

func good() {
	var x int
	func() {
		x = 2
	}()
}

func main() {
	body, err := Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
