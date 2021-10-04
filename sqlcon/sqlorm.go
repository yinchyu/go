package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Userlist struct {
	gorm.Model
	User string
	Age  int
	Sex  bool
}

type User struct {
	gorm.Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(www.yinchangyu.top:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	// db.AutoMigrate(new(User))
	// var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	// db.Create(&users)
	// batch insert from `[]map[string]interface{}{}`
	// 根据 map 创建记录时，association 不会被调用，且主键也不会自动填充
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
	user := &User{}
	// 关于？ 的使用类似 printf 的格式， 然后输出不同的数据类型
	res := db.Where("name = ?", "jinzhu_1").First(&user)

	fmt.Println(res.Error, res.RowsAffected)
	fmt.Println(user.CreatedAt, user.DeletedAt)
	times := user.CreatedAt
	// 默认填充的是时间的0值
	fmt.Println(times.Second())
}
