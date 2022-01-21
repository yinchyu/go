package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func IsLittleEndian() bool {
	// 对于大端和小端的理解也就到这里了
	var value int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)
	if *pb != 1 {
		return false
	}
	return true
}

// only impl by myself
func SwapEndianUin32(val uint32) uint32 {
	// 这个位移也是可以理解的， 每次操作一个字节
	return (val&0xff000000)>>24 | (val&0x00ff0000)>>8 |
		(val&0x0000ff00)<<8 | (val&0x000000ff)<<24
}

// use encoding/binary
// bigEndian littleEndian
func BigEndianAndLittleEndianByLibrary() {
	var value uint32 = 1
	by := make([]byte, 4)
	by2 := make([]byte, 4)
	binary.BigEndian.PutUint32(by, value)
	binary.LittleEndian.PutUint32(by2, value)

	fmt.Println(by, by2)
	fmt.Println("转换成大端后 ", by)
	fmt.Println("使用大端字节序输出结果：", binary.BigEndian.Uint32(by))
	little := binary.LittleEndian.Uint32(by)
	fmt.Println("大端字节序使用小端输出结果：", little)
}

func main() {
	BigEndianAndLittleEndianByLibrary()
	fmt.Println("当前系统是否为小端模式：", IsLittleEndian())
	fmt.Println("小端转换为大端后的结果：", SwapEndianUin32(1))
	fmt.Println(1 | 2)
}
