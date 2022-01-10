package main

import (
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type poolByte struct{ buf []byte } // 这种方式才能过语法检查

var (
	bytePool = sync.Pool{New: func() interface{} {
		return &poolByte{buf: make([]byte, 32<<10)}
	}}
	respOk = []byte("ok")
	//go:embed fileServer.ico
	icoData []byte // 嵌入图标文件
	//go:embed commandline.txt
	cli string
	//go:embed temp.html
	htmlTpl    string
	basePath   string // 传入路径的绝对路径
	useEncrypt string // 加密秘钥
	execPath   string // 可执行程序绝对路径

	unitByte = []struct {
		byte float64
		unit string
	}{
		{byte: 1},
		{byte: 1 << 10, unit: "B"},
		{byte: 1 << 20, unit: "KB"},
		{byte: 1 << 30, unit: "MB"},
		{byte: 1 << 40, unit: "GB"},
		{byte: 1 << 50, unit: "TB"},
	}
)

func InternalIp() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	ips := make([]string, 0, len(interfaces))
	for _, inf := range interfaces {
		if inf.Flags&net.FlagUp != net.FlagUp ||
			inf.Flags&net.FlagLoopback == net.FlagLoopback {
			continue
		}

		addr, err := inf.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addr {
			//  使用接口重新断言回去,方法就会增多
			if ipNet, ok := a.(*net.IPNet); ok &&
				!ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
func main() {
	var err error // 获取程序运行路径
	execPath, err = os.Executable()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 2 && os.Args[1] == "cli" {
		err = clientMain(os.Args[2:])
	} else {
		err = serverMain(os.Args[1:])
	}
	if err != nil {
		fmt.Println(err)
	}
}

func serverMain(args []string) error {
	// not specified name
	myFlag := flag.NewFlagSet(execPath, flag.ExitOnError)
	myFlag.StringVar(&basePath, "p", ".", "path")
	var addrStr string
	myFlag.StringVar(&addrStr, "s", ":9010", "ip:port")
	myFlag.StringVar(&useEncrypt, "e", "", "password")
	timeout := myFlag.Duration("t", time.Second*30, "server timeout")
	//reg := myFlag.Bool("reg", false, "add right click registry")
	_ = myFlag.Parse(args)

	tcpAddr, err := net.ResolveTCPAddr("tcp", addrStr)
	if err != nil {
		return err
	}

	addr, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	var urls []string
	addrStr = addr.Addr().String()
	if len(tcpAddr.IP) <= 0 {
		_, port, err := net.SplitHostPort(addrStr)
		if err != nil {
			return err
		}
		// 添加本机所有可用IP,组装Port
		if ips := InternalIp(); len(ips) > 0 {
			for _, v := range ips {
				urls = append(urls, v+":"+port)
			}
			addrStr = urls[0] // 取第一个IP作为默认url
		}
		urls = append(urls, "127.0.0.1:"+port) // 本地IP也可以
	} else {
		urls = []string{addrStr}
	}

	basePath, err = filepath.Abs(basePath)
	if err != nil {
		return err
	}

	//goland:noinspection HttpUrlsUsage

	tpl, err := template.New("").Parse(cli)
	if err != nil {
		return err
	}
	// 渲染命令行帮助
	err = tpl.Execute(os.Stdout, map[string]interface{}{
		"exec":    execPath,
		"addr":    addrStr,
		"dir":     basePath,
		"timeout": timeout.String(),
		"pass":    useEncrypt,
		"urls":    urls,
	})
	if err != nil {
		return err
	}

	if useEncrypt != "" {
		useEncrypt = md5str(useEncrypt)
	}
	http.HandleFunc("/", upDownFile)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) {
		// 不需要从请求中获取任何的信息
		_, _ = w.Write(icoData) // 网页的图标
	})
	return (&http.Server{ReadHeaderTimeout: *timeout}).Serve(addr)
}
func clientMain(args []string) error {
	myFlag := flag.NewFlagSet(execPath+" cli", flag.ExitOnError)
	data := myFlag.String("d", "", "<raw string> or @tmp.txt")
	output := myFlag.String("o", "", "output")
	point := myFlag.Bool("c", false, "Resumed transfer offset")
	timeout := myFlag.Duration("t", time.Minute, "client timeout")
	myFlag.StringVar(&useEncrypt, "e", "", "password")
	_ = myFlag.Parse(args)

	httpUrl := myFlag.Arg(0)
	if httpUrl == "" {
		return NewWebErr("url is null")
	}

	pool := bytePool.Get().(*poolByte)
	defer bytePool.Put(pool)

	clientHttp = &http.Client{Timeout: *timeout} // 使用加了超时的client
	key, c, err := testServer(httpUrl, pool.buf)
	if err != nil {
		return err
	}

	if *data != "" {
		return clientPost(*data, httpUrl, *point, key, c, pool.buf)
	}
	return clientGet(httpUrl, *output, *point, key, c, pool.buf)
}
