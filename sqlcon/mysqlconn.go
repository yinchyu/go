package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
	"time"
)

var conter = 0

type op struct {
	name string
	conn *sql.DB
}

func (op *op) connect() {
	dsn := "root:123456@tcp(www.yinchangyu.top:3306)/cov"
	// open 校验连接的地址格式是否正确
	dbconn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = dbconn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功")
	// 赋值给结构体对应的字段field
	op.conn = dbconn
}
func (op *op) exec() {
	rows, err := op.conn.Query(" select filter_user ,filter_time from  voice_op_record_cn")
	errorcheck(err)
	type per struct {
		Name string
		Time time.Time
	}
	var p per
	alldata := make(map[int]per)
	var v interface{}
	// 由于要事先定义类I型，所以对于获取到的结果需要先定义结构体去接受数据
	for rows.Next() {
		// 出现类型转换问题， null 不能转换为string
		err := rows.Scan(&p.Name, &v)
		// a:=reflect.TypeOf(v)
		// fmt.Println(p.name,a)
		if err != nil {
			fmt.Println(err, p.Name)
		}
		// 整形的切片可以转换为string 类型数据
		switch v.(type) {
		case []uint8:
			a := v.([]uint8)
			c, _ := time.Parse("2006-01-02 15:04:05", string(a))
			p.Time = c
			alldata[conter] = p
			conter += 1
		}
		fmt.Println("存入数据长度是：", len(alldata))

	}
	file, err := os.Create("./hand.ini")
	defer file.Close()
	wr := new([]byte)
	for k, zz := range alldata {
		z := []byte(strconv.Itoa(k))
		fmt.Println(string(z))
		*wr = append(*wr, '[')
		*wr = append(*wr, z...)
		*wr = append(*wr, ']', '\n')
		str1 := []byte("name=" + zz.Name)
		str2 := []byte("time=" + zz.Time.Format("2006-01-02 15:04:05"))
		*wr = append(*wr, str1...)
		*wr = append(*wr, '\n')
		*wr = append(*wr, str2...)
		*wr = append(*wr, '\n')
		n, err := file.Write(*wr)
		errorcheck(err)
		fmt.Println("write length n", n)

	}
}
func (op *op) ins() {
	res, err := op.conn.Exec(`insert into hotsearch (time ,id ,context) values("2020-01-09",12,"string")`)
	fmt.Println(res, err)
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

}
func (op *op) update() {
	sqlstynax := "update hotsearch set id=? where time=?"
	res, err := op.conn.Exec(sqlstynax, 123, "2020-01-09 00:00:00")
	fmt.Println(res, err)
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}
func (op *op) delete() {
	// 可以先将命令部分发送给mysql 端，然后发送数据 然后返回执行结果
	sqlstynax := "delete  from hotsearch where time=?"
	res, err := op.conn.Exec(sqlstynax, "2020-01-04 00:00:00")
	fmt.Println(res, err)
	//" 删除数据后最后一行的数据"
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}

func errorcheck(err error) {
	if err != nil {
		log.Fatal("错误信息:", err)
	}
}
func (op *op) prepare() {
	sqlstynax := "delete  from hotsearch where time=?"
	stmt, err := op.conn.Prepare(sqlstynax)
	errorcheck(err)
	res, err := stmt.Exec("2020-01-09 00:00:00")
	errorcheck(err)
	fmt.Println(res)
}
func prepare(tx *sql.Tx) {
	sqlstynax := "delete  from hotsearch where time=?"
	stmt, err := tx.Prepare(sqlstynax)
	errorcheck(err)
	res, err := stmt.Exec("2020-01-09 00:00:00")
	errorcheck(err)
	fmt.Println(res.RowsAffected())
}

var db *sqlx.DB

type user struct {
	Id   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

func initMySQL() (err error) {
	dsn := "root:123456@tcp(www.yinchangyu.top:3306)/cov"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	// db.SetMaxOpenConns(200)
	// db.SetMaxIdleConns(10)
	return
}

func sqlmulti() {
	var ops op
	ops.name = "first"
	ops.connect()
	// ops.delete()
	// 可以通过预编译命令的方式后续就可以只改变对应的参数就可以了
	// ops.exec()
	tx, err := ops.conn.Begin()
	errorcheck(err)
	// for i:=0;i<10;i++{
	// 	ops.ins()
	// }
	prepare(tx)
	// ops.prepare()
	err = tx.Commit()
	errorcheck(err)

	time.Sleep(time.Second * 6)
	err = tx.Rollback()
	errorcheck(err)

}
func used() {
	initMySQL()
	sqlsynatx := "select dt as a, id as b from  hotsearch where dt='%s'"
	// ' or 1=1 # sql 注入， 查询到所有的列
	sqlsynatx = fmt.Sprintf(sqlsynatx, "2020-03-16 21:47:26")
	fmt.Println(sqlsynatx)
	type per struct {
		Name string `db:"a"`
		// []uint8, 数据库中的时间提取有问题， 不能直接复制到时间对象上，然后[]int8 空的切片 不能转换 string
		// 结构体对于不同的包有不同的注释
		Time string `db:"b"`
		// missing destination name  表明结构体中的字段没有使用导出类型
		// missing destination name b in *[]main.per
	}
	// var p []per
	p := make([]per, 0)
	err := db.Select(&p, sqlsynatx)
	errorcheck(err)
	fmt.Printf(" 通过sqlx 来操作数据库 %v", p)
	// 使用sql 注入的关键是需要通过匹配前边的，然后屏蔽后边的
}
