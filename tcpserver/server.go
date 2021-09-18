package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"sync"
)

type broadcast struct {
	usermap map[string]*user
	sync.Mutex
}

type message struct {
	id      int64
	types   string
	content string
}

type user struct {
	userid          string
	conn            *net.Conn
	seq             int64
	messagesendchan chan message
	messagerecvchan chan message
}

var (
	userop = broadcast{
		usermap: make(map[string]*user),
	}
	counter = 1
	mux     sync.Mutex
)

func errorhandler(errortype string, err error) {
	if err != nil {
		fmt.Println(errortype, err)
	}
}

func listener() {
	listener, err := net.Listen("tcp", ":80")
	errorhandler("listen error", err)
	for {
		conn, err := listener.Accept()
		// 使用时间戳+连接地址
		uid := getuid()
		userop.Lock()
		userop.usermap[uid] = &user{
			userid: uid,
			conn:   &conn,
			seq:    0,
			// 暂时设置保存10条消息
			messagesendchan: make(chan message, 10),
			messagerecvchan: make(chan message, 10),
		}
		userop.Unlock()
		errorhandler("accept error", err)
		fmt.Println("send hello")
		userop.usermap[uid].firstconn()
		go userop.usermap[uid].handlerrecv()
		go userop.usermap[uid].handlersend()

	}
}

func (u *user) handlerrecv() {
	// 需要解决粘包的问题,当让不同的连接之间数据肯定是不会冲突的
	data := make([]byte, 1024)
	for {
		n, err := (*u.conn).Read(data)
		// scaner:=bufio.NewScanner(*u.conn)
		// for scaner.Scan(){
		// 	fmt.Println("read through scaner ",scaner.Text())
		// }
		errorhandler("read data err", err)
		if err != nil {
			// 断开连接之后直接返回
			u.colseconn()
			return
		}
		fmt.Println("read data length:", n, (*u.conn).RemoteAddr().String())
		msg := message{
			id:      u.seq + 1,
			types:   "broabcast",
			content: string(data),
		}
		u.messagesendchan <- msg
		u.seq++
	}
}

func (u *user) handlersend() {
	// 需要解决粘包的问题,当让不同的连接之间数据肯定是不会冲突的
	for {
		select {
		case msg, ok := <-u.messagesendchan:
			if ok {
				userop.Lock()
				for id, uu := range userop.usermap {
					if id != u.userid {
						(*uu.conn).Write([]byte(u.userid + ": " + msg.content))
					}
				}
				userop.Unlock()
			} else {
				// channel 已经进行了关闭,
				return
			}

		}
	}
}
func (u *user) firstconn() {
	for id, uu := range userop.usermap {
		if id == u.userid {
			fmt.Fprint(*uu.conn, "welcome")
		} else {
			fmt.Fprint(*uu.conn, u.userid+": "+"online")
		}
	}
}
func (u *user) colseconn() {
	for id, uu := range userop.usermap {
		if id == u.userid {
			// 表明是自己,就不用写了
			(*u.conn).Close()
			close(u.messagesendchan)
			close(u.messagerecvchan)
		} else {
			fmt.Fprint(*uu.conn, u.userid+": "+"offline")
		}
	}
}

func getuid() string {
	mux.Lock()
	defer mux.Unlock()
	counter++
	return strconv.Itoa(counter)
}

func main() {
	listener()
}
