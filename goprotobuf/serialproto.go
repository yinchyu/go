package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"os"
)

func main() {
	stu := &StudentList_Student{
		Name:   "geek",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	stulist := &StudentList{
		TargetList: []*StudentList_Student{},
	}
	//  添加了三个组来进行操作
	stulist.TargetList = append(stulist.TargetList, stu)
	stulist.TargetList = append(stulist.TargetList, stu)
	stulist.TargetList = append(stulist.TargetList, stu)
	data, err := proto.Marshal(stu)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	// data1, err1 := proto.Marshal(stulist)
	// if err1 != nil {
	// 	log.Fatal("marshaling error: ", err)
	// }
	// data 写入到文件中然后再通过文件进行 读取
	// file,err:=os.Create("dump_file")
	// file.Write(data1)
	// file.Close()
	file2, err2 := os.Open("newdump")
	if err2 != nil {
		log.Println(err2)
	}
	data3, err3 := io.ReadAll(file2)
	if err3 != nil {
		log.Println(err3)
	}
	file2.Close()
	fmt.Println("====", len(data3))
	//解析出来的是uint8类型的二进制文件， 然后通过 unmarshal进行反序列化操作，就可以进行还原
	fmt.Printf("start ,%v %v %T\n", data, string(data), data)
	newstu := &StudentList_Student{}
	newstulist := &StudentList{}
	err = proto.Unmarshal(data, newstu)
	err = proto.Unmarshal(data3, newstulist)
	// repeated 限定字段表明这个数据类型是列表类型的数据
	fmt.Printf("%v %T \n", newstu.Scores, newstu.Scores)
	fmt.Printf("%v %T \n", newstu.Name, newstu.Name)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(len(newstulist.GetTargetList()))
	for _, v := range newstulist.GetTargetList() {
		fmt.Println(v.GetName())
	}
	// 会把内部的所有的元素都给清楚掉
	newstulist.Reset()
	fmt.Println(len(newstulist.GetTargetList()))
	for _, v := range newstulist.GetTargetList() {
		fmt.Println(v.GetName())
	}
}
