package main

import (
	"log"
)

var (
	mydb, err = InitDB()
)

func main() {
	if err != nil {
		log.Println(err)
	}
	r := router()
	r.Run(":8080")
	// data:=mydb.FindGoodById(1197)
	// // 差的时间通过database/sql结构体来描述合法性
	// fmt.Println(data.LastUpdateTime.Valid)
	// fmt.Println(data)
	// datalist:=mydb.FindParamsByGoodsId(1197)
	// // 差的时间通过database/sql结构体来描述合法性
	// for i:=range datalist{
	// 	fmt.Println(datalist[i].GoodsId)
	//
	// }

}
