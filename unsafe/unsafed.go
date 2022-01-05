package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"unsafe"
)

type User1 struct {
	A int32
	B []int32
	C string
	D bool
}

type User2 struct {
	B []int32
	A int32
	D bool
	C string
}

type User3 struct {
	D bool
	B []int32
	A int32
	C string
}

// struct{}放到后边容易占用空间
type test1 struct {
	a int32
	b struct{}
}

// 放到前边不占用空间
type test2 struct {
	a struct{}
	b int32
}

func testfunc() {

	fmt.Println(unsafe.Sizeof(test1{})) // 8
	fmt.Println(unsafe.Sizeof(test2{})) // 4
	var u1 User1
	var u2 User2
	var u3 User3
	fmt.Println("string size is ", unsafe.Sizeof(""))
	fmt.Println("u1 size is ", unsafe.Sizeof(u1))
	fmt.Println("u2 size is ", unsafe.Sizeof(u2))
	fmt.Println("u3 size is ", unsafe.Sizeof(u3))
	func_example()
}
func func1() {
	CopyField()
	type Human struct {
		sex  bool
		age  uint8
		min  int
		name string
	}

	h := Human{
		true,
		30,
		1,
		"hello",
	}
	i := unsafe.Sizeof(h)
	j := unsafe.Alignof(h.age)
	k := unsafe.Offsetof(h.min)
	m := unsafe.Offsetof(h.name)
	fmt.Println(i, j, k, m)
	fmt.Printf("%p\n", &h)
	var p unsafe.Pointer
	p = unsafe.Pointer(&h)
	fmt.Println(p)

}

type User struct {
	Name   string
	Age    uint32
	Gender bool // 男:true 女：false 就是举个例子别吐槽我这么用。。。。
}

func func_example() {
	// sizeof
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int(10)))
	fmt.Println(unsafe.Sizeof(int32(190)))
	fmt.Println(unsafe.Sizeof("asong"))        // 16
	fmt.Println(unsafe.Sizeof([]int{1, 3, 4})) //24
	// Offsetof
	user := User{Name: "Asong", Age: 23, Gender: true}
	userNamePointer := unsafe.Pointer(&user)

	nNamePointer := (*string)(unsafe.Pointer(userNamePointer))
	*nNamePointer = "Golang梦工厂"
	p1 := uintptr(userNamePointer)
	nAgePointer := (*uint32)(unsafe.Pointer(p1 + unsafe.Offsetof(user.Age)))
	*nAgePointer = 25

	nGender := (*bool)(unsafe.Pointer(uintptr(userNamePointer) + unsafe.Offsetof(user.Gender)))
	*nGender = false

	fmt.Printf("u.Name: %s, u.Age: %d,  u.Gender: %v\n", user.Name, user.Age, user.Gender)
	// Alignof
	var b bool //1
	var i8 int8
	var i16 int16
	var i64 int64
	var f32 float32
	var s string // 8
	var m map[string]string
	var p *int32   // 8 32 + 16 = 48 + 1 = 49 + 7
	var p2 []int32 //8
	var p3 int32
	//string 是8,一个slice 是24
	fmt.Println(unsafe.Alignof(b))
	fmt.Println(unsafe.Alignof(i8))
	fmt.Println(unsafe.Alignof(i16))
	fmt.Println(unsafe.Alignof(i64))
	fmt.Println(unsafe.Alignof(f32))
	fmt.Println(unsafe.Alignof(s))
	fmt.Println(unsafe.Alignof(m))
	fmt.Println(unsafe.Alignof(p))
	fmt.Println(unsafe.Alignof(p2))
	fmt.Println(unsafe.Alignof(p3))
}

func example_one() {
	number := 5
	pointer := &number
	fmt.Printf("number:addr:%p, value:%d\n", pointer, *pointer)

	float32Number := (*float32)(unsafe.Pointer(pointer))
	*float32Number = *float32Number + 3

	fmt.Printf("float64:addr:%p, value:%f\n", float32Number, *float32Number)
}

func stringToByte(s string) []byte {
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	newHeader := reflect.SliceHeader{
		Data: header.Data,
		Len:  header.Len,
		Cap:  header.Len,
	}
	// ()() 两个括号少不了
	// header 转换成为unsafe.pointer ，然后进行强制类型转换
	return *(*[]byte)(unsafe.Pointer(&newHeader))
}

func bytesToString(b []byte) string {
	// 转换成通用指针，底层的内存布局是相同的， 就可以进行转换
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	newHeader := reflect.StringHeader{
		Data: header.Data,
		Len:  header.Len,
	}

	return *(*string)(unsafe.Pointer(&newHeader))
}

func lbytesToString(b []byte) string {

	return *(*string)(unsafe.Pointer(&b))
}

type My struct {
	a string
	b int
}

// 修改一个对应的结构体的值通过指针进行修改
func CopyField() {
	var w *My = new(My)
	fmt.Println(w.a, w.b)
	baddr := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*(*int)(baddr) = 10
	fmt.Println(w.a, w.b)
}
func ReadBody() {
	response, err := http.Get("https://go.dev/")

	if err != nil {
		log.Println(err)
		return
	}
	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(bytesToString(readAll))
	defer response.Body.Close()
}

func AddPointer() {
	var a int64
	var b float64
	b = 12.23
	a = *(*int64)(unsafe.Pointer(&b))
	fmt.Println(a)
}
func ChangeStructfield() {
	type user struct {
		name    string
		age     int
		company string
	}
	u := new(user)  // A
	fmt.Println(*u) // { 0}

	uName := (*string)(unsafe.Pointer(u)) // B
	*uName = "Jemmy"
	fmt.Println(*u) // {Jemmy 0}

	uAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age))) // C
	*uAge = 23
	fmt.Println(*u) // {Jemmy 23}

	uCompany := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.company))) // D
	*uCompany = "吹牛逼技术有限公司"
	fmt.Println(*u) // {Jemmy 23 吹牛逼技术有限公司}
}

func main() {
	AddPointer()
	ChangeStructfield()
}
