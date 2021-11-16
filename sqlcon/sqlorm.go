package main

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
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

type Student struct {
	gorm.Model
	Name     string
	ClassId  int
	Classs    Class `gorm:"foreignkey:ClassName;reference:ClassId"`
	Teachers []Teacher `gorm:"many2many:student_teacher"`
}
type Class struct {
	gorm.Model
	ClassId int
	ClassName string
}



type Teacher struct {
	gorm.Model
	TeacherName string
	classId int
	TeachClass  Class `gorm:"foreignkey:ClassName;reference:ClassId"`
	Students  []Student
}


// User 拥有并属于多种 language，`user_languages` 是连接表
type Names struct {
	gorm.Model
	Name         string
	// 这个表中有令一个表的主键,作为这个表的外键来操作， 才可以进行处理操作
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
	// 使用 CompanyRefer 作为外键
}

type Company struct {
	// 生成表格之后对应的属性可能不会进行更改
	ID   int `gorm:"type:int"`
	Name string
	Age int
	Read int
}

func AddUser(db *gorm.DB){
	db.AutoMigrate(new(User))


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
//func main() {
//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
//	dsn := "root:123456@tcp(www.yinchangyu.top:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Println(err)
//	}
//	db.AutoMigrate(new(Names))
//
//}
