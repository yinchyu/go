package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Name struct {
	Names string `gorm:"<-:create"` // 允许读和创建
	//Name1 string `gorm:"<-:update"` // 允许读和更新
	//Name2 string `gorm:"<-"`        // 允许读和写（创建和更新）
	//Name3 string `gorm:"<-:false"`  // 允许读，禁止写
	//Name4 string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
	//Name5 string `gorm:"->;<-:create"` // 允许读和写
	//Name6 string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
	//Name7 string `gorm:"-"`  // 通过 struct 读写会忽略该字段
}
func main() {
	dsn:="root:123456@tcp(www.yinchangyu.top:3306)/gorm?charset=utf8mb4&parseTime=true&loc=Local"
	dbconn,err:=gorm.Open(mysql.Open(dsn))
	if err!=nil{
		log.Println(err)
	}
	dbconn.AutoMigrate(new(Name))
	// 设置只读属性实在框架的层面做的， 然后对应的语句应该不会被执行
	dbconn.Create(&Name{Names: "1024"})
	dbconn.Where("names").Updates(&Name{Names: "2048"})
	res:=Name{}
	resdb:=dbconn.Where("names=?",1024).First(&res)
	fmt.Println("打印出错误的位置信息:",resdb.Error,resdb.RowsAffected)
	fmt.Println(res)

}
