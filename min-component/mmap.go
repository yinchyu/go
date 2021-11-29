package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const defaultMaxFileSize = 1 << 30        // 假设文件最大为 1G
// 映射的大小可以远远小于文件的大小， 因为会进行换页的操作
const defaultMemMapSize = 128 * (1 << 20) // 假设映射的内存大小为 128M

type Demo struct {
	file    *os.File
	data    *[defaultMaxFileSize]byte
	dataRef []byte
}

func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(msg, v...))
	}
}
func (demo *Demo) mmap() {
	h, err := syscall.CreateFileMapping(syscall.Handle(demo.file.Fd()), nil, syscall.PAGE_READWRITE, 0, defaultMemMapSize, nil)
	_assert(h != 0, "failed to map", err)

	addr, err := syscall.MapViewOfFile(h, syscall.FILE_MAP_WRITE, 0, 0, uintptr(defaultMemMapSize))
	_assert(addr != 0, "MapViewOfFile failed", err)

	err = syscall.CloseHandle(syscall.Handle(h))
	_assert(err == nil, "CloseHandle failed")

	// Convert to a byte array.
	demo.data = (*[defaultMaxFileSize]byte)(unsafe.Pointer(addr))
}

func (demo *Demo) munmap() {
	addr := (uintptr)(unsafe.Pointer(&demo.data[0]))
	_assert(syscall.UnmapViewOfFile(addr) == nil, "failed to munmap")
}

func Mmapwirte() {
	os.Chdir("D:\\桌面文件夹\\gotest\\test2")
	_ = os.Remove("tmp.txt")
	f, _ := os.OpenFile("tmp.txt", os.O_CREATE|os.O_RDWR, 0644)
	demo := &Demo{file: f}
	demo.mmap()
	// 解除之间的关联
	defer demo.munmap()

	msg := "hello geektutu!"
	for i, v := range msg {
		demo.data[2*i] = byte(v)
		demo.data[2*i+1] = byte(' ')
	}
}
