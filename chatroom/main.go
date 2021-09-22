package main

import (
	"fmt"
	"github.com/polaris1119/chatroom/global"
	"github.com/polaris1119/chatroom/logic"
	"github.com/polaris1119/chatroom/server"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var (
	addr   = ":80"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

API Listen on：%s
`
)

func init() {
	global.Init()
	// 初始化完成之后有三个变量
	//rootdir  sensitiveword queuelen
}

func main() {
	fmt.Printf(banner, addr)
	// 广播消息处理
	go logic.Broadcaster.Start()
	// 三个handle 处理函数
	http.HandleFunc("/", server.HomeHandleFunc)
	http.HandleFunc("/user_list", server.UserListHandleFunc)
	http.HandleFunc("/ws", server.WebSocketHandleFunc)
	log.Fatal(http.ListenAndServe(addr, nil))
}
