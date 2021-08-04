package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Client struct {
	Name string
}
 type   LoadBalancer struct {
 	client []*Client
 	size int32
 }

 func NewLoadBalancer(size int32) *LoadBalancer{
 	clientlist:=make([]*Client,size)
 	for i:=0;i<int(size);i++{
 		clientlist[i]=&Client{strconv.Itoa(i)}
	}
 	return &LoadBalancer{
 		clientlist,
 		size}
 }

func ( m *LoadBalancer)GetClient() *Client {
	rand.Seed(time.Now().Unix())
	x :=rand.Int31n(100)
	return m.client[x%m.size]
}

func (m *Client) Do() {
	fmt.Println("do", m.Name)

}

func main(){
	loader:=NewLoadBalancer(5)
	// 随机获取一个
	loader.GetClient().Do()


}