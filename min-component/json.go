package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	t := struct {
		N int
		time.Time
	}{
		5,
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
	}
	// 如果需要输出对应的正确就是需要重写   MarshalJson  方法
	m, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(t.MarshalJSON())
	fmt.Println(string(m))

}
