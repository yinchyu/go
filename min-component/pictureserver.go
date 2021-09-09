package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

// go:embed wechat.png
var picture string

func writelog(ctx context.Context, log chan string, filename string) {
	fs, err := os.Create(filename)
	defer fs.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case strs := <-log:
			fmt.Println("=======")
			_, err := fs.Write([]byte(strs))
			if err != nil {
				fmt.Println(err)
			}
		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}

}
func stop(cancal func()) {
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())
	cancal()
}

func main() {
	logchan := make(chan string, 1024)
	logname := "./log.txt"
	// context 必须要通过 withcancel 进行包装才能进行操作，
	// 通过包装来 withcancel, withtimeout, withvalue , withdeadline, 四个不同的参数来进行操作
	ctx, cancel := context.WithCancel(context.Background())
	go writelog(ctx, logchan, logname)
	conn, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)
	}
	hander := func(w http.ResponseWriter, user *http.Request) {
		wirtedata := user.Host + ": " + user.RequestURI + "\n"
		logchan <- wirtedata
		fmt.Print(wirtedata)
		io.WriteString(w, picture)
	}
	http.HandleFunc("/", hander)
	err = http.Serve(conn, nil)
	if err != nil {
		fmt.Println(err)
		// 关闭另一个协程， 避免协程的泄漏
		cancel()
	}
}
