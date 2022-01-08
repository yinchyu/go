package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type userinfo struct {
	Name     string `json:"user" form:"user" binding:"Require"`
	Password string `json:"password" form:"password"`
}

func Require(f1 validator.FieldLevel) bool {
	a := f1.Field().Interface().(string)
	if a == "ycy" {
		return true
	}
	return false
}

// 可以给不同的路由分组使用， 使用use 函数， 传递的参数可以是一个或者多个
// 使用链式调用， 或者使用递归调用
func middle1(c *gin.Context) {
	fmt.Println("在方法前调用")
	c.Next()
	fmt.Println("在方法后调用")
}
func middle2(c *gin.Context) {
	fmt.Println("在方法前调用")
	c.Next()
	fmt.Println("在方法后调用")
}
func demo1() {
	getfunc := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}

	router := gin.Default() //初始化一个gin实例,并且默认配置了Logger,Recovery
	// 通过关键字进行分组操作，就像是文件夹的操作
	v1 := router.Group("v1").Use(middle1, middle2)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("Require", Require)
	}
	v1.GET("/:id", func(c *gin.Context) {
		// localhost/gets/1234?name=ycy&password=1234, 传递的参数可以被解析 query 解析
		// 占位符传递参数
		id := c.Param("id")
		fmt.Println(id)
		// query 传递参数
		name := c.Query("name")
		password := c.Query("password")
		c.JSON(200, gin.H{
			"message":  "pong",
			"id":       id,
			"name":     name,
			"password": password,
		})
	})
	v1.POST("/", func(c *gin.Context) {
		// 直接使用postfrom 对所有的数据进行解析
		// name:=c.PostForm("name")
		// password:=c.PostForm("password")
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// 	"name":name,
		// 	"password":password,)}

		fname, _ := c.FormFile("filename")
		c.SaveUploadedFile(fname, "./"+fname.Filename)
		// 加上头部的描述信息
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fname.Filename))
		c.File("./" + fname.Filename)
		// 一种方式是通过  upload 上传文件
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// 	"data":fname,
		// })
	})
	v1.PUT("/", getfunc)
	v1.DELETE("/", getfunc)
	//接口路由，如果url不是/debug/vars，则用metricBeat去获取会出问题
	router.Run(":80")
}

//  通过header请求处理数据
func ginHeader() {
	engine := gin.Default()
	groupv1 := engine.Group("v1")
	groupv1.HEAD("/header", func(c *gin.Context) {
		fmt.Println("enter header")

		getval := c.GetHeader("Content-type")
		fmt.Println(getval)

	})
	engine.Run(":8090")
}

// 通过http的请求来处理数据， gin 将一个不同的
func httpHeader(){
	hand := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("enter")
		switch r.Method {
		case http.MethodHead:
			getval := r.Header.Get("Content-Type")
			fmt.Println(getval)
		case http.MethodGet:
			data := []byte("hello guest")
			w.Write(data)
			fmt.Println(string(data))
		case http.MethodPost:
			// 对应的from 格式必须先进行 parsefrom 后才能在get 中得到数据
			fmt.Println(r.ParseForm())
			fmt.Println(r.Form.Get("name"))

		}
	}
	http.HandleFunc("/v1/header", hand)
	http.ListenAndServe(":8090", nil)
}