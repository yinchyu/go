package main
import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)


func  connread(webconn *websocket.Conn){
	defer  webconn.Close()
	for{
		n,data,err:=webconn.ReadMessage()
		if err!=nil{
			log.Println(err)
			return
		}
		fmt.Println(n,string(data))
		connwrite(webconn)
	}
}
func connwrite(webconn *websocket.Conn){
		err:=webconn.WriteMessage(websocket.TextMessage,[]byte("helloworld"))
		if err!=nil{
			log.Println(err)
			return
		}
		time.Sleep(time.Second*1)

}


func server(){
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		upgrader:=websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		newconn,err:=upgrader.Upgrade(w,req,http.Header{})
		if err!=nil{
			log.Println(err)
		}
		go connread(newconn)
	}
	router:=http.NewServeMux()
	router.HandleFunc("/ws", helloHandler)
	// 使用默认的路由函数
	log.Fatal(http.ListenAndServe(":2021", router))
}
func main() {
	server()
}
