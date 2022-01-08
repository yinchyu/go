package main

import (
	"bytes"
	"crypto/md5"
	_ "embed"
	"fmt"
	"net"
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

func InternalIp() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	ips := make([]string, 0, len(interfaces))
	for _, inf := range interfaces {
		// 可以直接看到一个接口是否启动了
		if inf.Flags&net.FlagUp != net.FlagUp ||
			inf.Flags&net.FlagLoopback == net.FlagLoopback {
			continue
		}

		addr, err := inf.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addr {
			if ipNet, ok := a.(*net.IPNet); ok &&
				!ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

func IP() {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for i := range interfaces {
		addrs, err := interfaces[i].Addrs()
		if err != nil {
			return
		}
		for j := range addrs {
			//  使用接口重新断言回去
			if ip, ok := addrs[j].(*net.IPNet); ok {
				fmt.Println(ip)
			}

		}
	}
}

//go:embed sum.s
var data []byte

func md5s() {
	hash := md5.New()
	hash.Write(data)
	fmt.Printf("%x\n", hash.Sum(nil))
	hash.Reset()
	fmt.Printf("%x\n", hash.Sum(nil))
}
