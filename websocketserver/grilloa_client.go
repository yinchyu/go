package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func InitScoket() *websocket.Conn{
	// 请求头什么信息都不添加
	conn,_,err:=websocket.DefaultDialer.Dial("ws://localhost:2021/ws",http.Header{})
	if err!=nil{
		log.Println(err)
	}
	return conn
}


func ReceiveMessage(conn *websocket.Conn){
	for  {
		_,data,err:=conn.ReadMessage()
		if err!=nil{
			log.Println("read error:",err)
		}
		fmt.Println(string(data))
		SendMessage(conn, data)
	}
}

func SendMessage(conn *websocket.Conn, data []byte){
err:=conn.WriteMessage(websocket.TextMessage,data)
if err!=nil{
	log.Println("write error:",err)
	return
}
}


func  SendAllow(conn *websocket.Conn , chancel context.Context){
	timmer:=time.NewTicker(time.Second*2)
	defer timmer.Stop()
	for {
		select {
		case ch:= <-timmer.C:
			fmt.Println("send data ",ch)
			SendMessage(conn ,[]byte(ch.String()))
		case <-chancel.Done():
			fmt.Println("time send will be cancel")
			return
		}
	}
}

func main() {
conn :=InitScoket()
ctx,cancel:=context.WithTimeout(context.Background(),time.Second*10)
// 确保最后也会被关掉
defer cancel()
defer conn.Close()
go SendAllow(conn, ctx)
ReceiveMessage(conn)
}
