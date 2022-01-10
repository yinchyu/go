package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func scanSize(h http.Header) (first, last, length int64, err error) {
	var n int // Content-Range: bytes (unit first byte pos) - [last byte pos]/[entity length]
	n, err = fmt.Sscanf(h.Get("Content-Range"), "bytes %d-%d/%d", &first, &last, &length)
	if n != 3 {
		err = fmt.Errorf("scanSize n=%d", n)
	}
	return
}

func parseInt(s string) (int64, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err // 此时必须返回0,针对忽略err的场景
	}
	return n, nil
}
