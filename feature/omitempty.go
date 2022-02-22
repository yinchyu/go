package main

import (
	"encoding/json"
	"fmt"
)

type address struct {
	Street     string      `json:"street"`
	Ste        string      `json:"suite,omitempty"`
	City       string      `json:"city"`
	State      string      `json:"state"`
	Zipcode    string      `json:"zipcode"`
	Coordinate *coordinate `json:"coordinate,omitempty"`
}

type coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

// 如果 omitempty 如果原来本身是空的话就直接忽略，也不会进行输出
// 嵌套的结构体不受这个的限制，必须将嵌套的结构体改为指针类型的才可以
func main() {
	data := `{
  "street": "200 Larkin St",
  "suite":"",
  "city": "San Francisco",
  "state": "CA",
  "zipcode": "94102"
 }`
	addr := new(address)
	json.Unmarshal([]byte(data), &addr)

	// 处理了一番 addr 变量...

	addressBytes, _ := json.MarshalIndent(addr, "", "    ")
	fmt.Printf("%s\n", string(addressBytes))
}
