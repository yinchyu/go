package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func add(a, b int) int // 汇编函数声明

func sub(a, b int) int // 汇编函数声明

func mul(a, b int) int // 汇编函数声明
func sum([]int64) int64

func call1() {
	fmt.Println(add(10, 11))
	fmt.Println(sub(99, 15))
	fmt.Println(mul(11, 12))
}

var offsetDict = map[string]int64{
	"go1.4":  128,
	"go1.5":  184,
	"go1.6":  192,
	"go1.7":  192,
	"go1.8":  192,
	"go1.9":  152,
	"go1.10": 152,
	"go1.11": 152,
	"go1.12": 152,
	"go1.13": 152,
	"go1.14": 152,
	"go1.17": 152,
}

var offset = func() int64 {
	ver := strings.Join(strings.Split(runtime.Version(), ".")[:2], ".")
	return offsetDict[ver]
}()

// GetGoID returns the goroutine id
func GetGoID() int64

func ExtractGID(s []byte) int64 {
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

// Parse the goid from runtime.Stack() output. Slow, but it works.
func getSlow() int64 {
	var buf [64]byte
	return ExtractGID(buf[:runtime.Stack(buf[:], false)])
}
func Goid() {
	var wg sync.WaitGroup
	wg.Add(1)
	id := GetGoID()
	go func() {
		idInternal := GetGoID()
		fmt.Println(id, idInternal)
		wg.Done()
	}()
	wg.Wait()
}
func main() {

}
