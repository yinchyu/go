package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)
func  nhooyr_server(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"muxter router")
	})
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn,err:=websocket.Accept(writer,request,nil)
		if err!=nil{
			log.Println(err)
		return
		}
		defer conn.Close(websocket.StatusInternalError,"internal error")
		ctx,cancel:=context.WithTimeout(request.Context(),time.Second*10)
		defer cancel()
		var v interface{}

		err=wsjson.Read(ctx,conn,&v)
		if err!=nil{
			log.Println(err)
			return
		}

		log.Printf("接收到客户端：%v\n", v)
		err = wsjson.Write(ctx, conn, "Hello WebSocket Client")
		if err != nil {
			log.Println(err)
			return
		}
		conn.Close(websocket.StatusNormalClosure, "")
	})
	log.Fatal(http.ListenAndServe(":9998", nil))
}

func main() {
	nhooyr_server()
}
