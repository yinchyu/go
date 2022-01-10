package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var request = `{"id":7044144249855934983,"name":"demo"}`

	var test interface{}
	err := json.Unmarshal([]byte(request), &test)
	if err != nil {
		fmt.Println("error:", err)
	}
	// 这种断言的方式是直接将
	obj := test.(map[string]interface{})

	dealStr, err := json.Marshal(test)
	if err != nil {
		fmt.Println("error:", err)
	}
	id := obj["id"]
	// 反序列化之后重新序列化打印,发现精度不一致，线上在所有的位置使用整数就会导致这个问题。
	fmt.Println(request)
	fmt.Println(string(dealStr))
	fmt.Printf("%+v\n", reflect.TypeOf(id).Name())
	fmt.Printf("%+v\n", id.(float64))
	decoder := json.NewDecoder(strings.NewReader(request))
	// 使用哦个number  不是使用float64 来解析
	decoder.UseNumber()
	var test2 interface{}
	decoder.Decode(&test2)
	objStr, err := json.Marshal(test2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(request)
	fmt.Println(string(objStr))
}
