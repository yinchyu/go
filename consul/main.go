package main

import (
	"github.com/hashicorp/consul/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)
import "fmt"

func main() {
	// Get a new client
	// 创建的server 在哪里， 如何进行自动的接收数据
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	// 设置对应的配置kv 的形式， 另一个机器应该也可以获取
	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	// Lookup the pair
	//pair, _, err = kv.Get("goready", nil)
	//if err != nil {
	//	panic(err)
	//}
	// 直接通过%s 进行转换
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	//创建一个新服务。
	registration := new(api.AgentServiceRegistration)
	registration.ID = "12345"
	registration.Name = "user-tomcat"
	registration.Port = 8080
	registration.Tags = []string{"user-tomcat"}
	registration.Address = "127.0.0.1"
	//增加check。
	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"
	//注册check服务。
	registration.Check = check
	log.Println("get check.HTTP:", check)

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	//err = client.Agent().ServiceDeregister("helloworld")
	//if err != nil {
	//	log.Fatal("register server error : ", err)
	//}
	entry := api.ExportedServicesConfigEntry{
		Name:        "",
		Partition:   "",
		Services:    nil,
		Meta:        nil,
		CreateIndex: 0,
		ModifyIndex: 0,
	}

	json, err := entry.MarshalJSON()
	if err != nil {
		return
	}
	fmt.Println("the json data ", string(json))

	errChan := make(chan error, 1)
	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		// 优雅退出
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	log.Println(error)

}
