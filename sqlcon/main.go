package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mysql/model"
	"os"
)

type Name struct {
	// 设置一个字段的权限
	gorm.Model
	Names string `gorm:"<-:create"` // 允许读和创建
	//Name1 string `gorm:"<-:update"` // 允许读和更新
	//Name2 string `gorm:"<-"`        // 允许读和写（创建和更新）
	//Name3 string `gorm:"<-:false"`  // 允许读，禁止写
	//Name4 string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
	//Name5 string `gorm:"->;<-:create"` // 允许读和写
	//Name6 string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
	//Name7 string `gorm:"-"`  // 通过 struct 读写会忽略该字段
}

func Readconfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("mysql.databasename"))
}

type Reader struct {
	// 所有的导出字段必须大写才能进行操作，只有大写之后才不会报错
	gorm.Model
	Name   string `gorm:"type:varchar(255)" `
	Reader string `gorm:"type:varchar(255)" `
}

func main() {
	dsn := "root:ycy1234@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(new(Name))
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Video{})
	db.AutoMigrate(&model.Review{})
	db.AutoMigrate(&model.Interactive{})
	db.AutoMigrate(&model.Follow{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Reply{})
	db.AutoMigrate(&model.Announce{})
	db.AutoMigrate(&model.AnnounceUser{})
	db.AutoMigrate(&model.Message{})
	db.AutoMigrate(&model.Danmaku{})
	db.AutoMigrate(&model.Carousel{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Collection{})
	db.AutoMigrate(&model.VideoCollection{})
	if err != nil {
		log.Println(err)
	}

	//设置只读属性实在框架的层面做的， 然后对应的语句应该不会被执行
	//db.Create(&Name{Names: "1024"})
	//db.Create(&Name{Names: "1024"})
	//db.Where("names").Updates(&Name{Names: "2048"})
	//res := Name{}
	//resdb := db.Where("names=?", 1024).First(&res)
	//fmt.Println(resdb.RowsAffected)

}
