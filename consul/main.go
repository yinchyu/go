package main

import (
	"encoding/json"
	"fmt"
)

//func main() {
//	// Get a new client
//	// 创建的server 在哪里， 如何进行自动的接收数据
//	client, err := api.NewClient(api.DefaultConfig())
//	if err != nil {
//		panic(err)
//	}
//
//	// Get a handle to the KV API
//	kv := client.KV()
//
//	// PUT a new KV pair
//	// 设置对应的配置kv 的形式， 另一个机器应该也可以获取
//	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
//	_, err = kv.Put(p, nil)
//	if err != nil {
//		panic(err)
//	}
//
//	// Lookup the pair
//	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
//	// Lookup the pair
//	//pair, _, err = kv.Get("goready", nil)
//	//if err != nil {
//	//	panic(err)
//	//}
//	// 直接通过%s 进行转换
//	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
//	//创建一个新服务。
//	registration := new(api.AgentServiceRegistration)
//	registration.ID = "12345"
//	registration.Name = "user-tomcat"
//	registration.Port = 8080
//	registration.Tags = []string{"user-tomcat"}
//	registration.Address = "127.0.0.1"
//	//增加check
//	check := new(api.AgentServiceCheck)
//	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
//	//设置超时 5s。
//	check.Timeout = "5s"
//	//设置间隔 5s。
//	check.Interval = "5s"
//	//注册check服务。
//	registration.Check = check
//	log.Println("get check.HTTP:", check)
//
//	err = client.Agent().ServiceRegister(registration)
//
//	if err != nil {
//		log.Fatal("register server error : ", err)
//	}
//
//	//err = client.Agent().ServiceDeregister("helloworld")
//	//if err != nil {
//	//	log.Fatal("register server error : ", err)
//	//}
//	entry := api.ExportedServicesConfigEntry{
//		Name:        "",
//		Partition:   "",
//		Services:    nil,
//		Meta:        nil,
//		CreateIndex: 0,
//		ModifyIndex: 0,
//	}
//
//	json, err := entry.MarshalJSON()
//	if err != nil {
//		return
//	}
//	fmt.Println("the json data ", string(json))
//
//	errChan := make(chan error, 1)
//	go func() {
//		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
//		// 优雅退出
//		c := make(chan os.Signal, 1)
//		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
//		errChan <- fmt.Errorf("%s", <-c)
//	}()
//
//	error := <-errChan
//	//服务退出取消注册
//	log.Printlmn(error)
//
//}
//func main() {
//	// 直接对应regionconfig
//	s := `{
//"RegionConfig":{
//"ID":{
//"ReviewTags":{
//"default":["ecom_write_review_tag_value_of_money","ecom_write_review_tag_compared_to_description","ecom_write_review_tag_quality_of_product"],
//"product_type_clothes_id":["ecom_write_review_tag_fabric_material","ecom_write_review_tag_value_of_money","ecom_write_review_tag_compared_to_description","ecom_write_review_tag_quality_of_product"]
//}
//},
//"default":{
//"ReviewTags":{
//"default":["ecom_write_review_tag_value_of_money","ecom_write_review_tag_compared_to_description","ecom_write_review_tag_quality_of_product"],
//"product_type_clothes_id":["ecom_write_review_tag_fabric_material","ecom_write_review_tag_value_of_money","ecom_write_review_tag_compared_to_description","ecom_write_review_tag_quality_of_product"]
//
//
//}
//}
//}
//}`
//
//	var h ReviewTagsInfo
//	var c T
//	var name map[string]interface{}
//	err := json.Unmarshal([]byte(s), &h)
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = json.Unmarshal([]byte(s), &c)
//	if err != nil {
//		fmt.Println(err)
//	}
//	// 当然也可以使用interface 包裹一切，只不过需要无数的断言
//	//interface{} 不能断言为  []stirng, 只能断言为[]interface然后断言为int类型数据
//	err = json.Unmarshal([]byte(s), &name)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("===", name["RegionConfig"].(map[string]interface{})["ID"].(map[string]interface{})["ReviewTags"].(map[string]interface{})["default"].([]interface{})[0].(string))
//	fmt.Println(h.RegionConfig["default"].ReviewTags["product_type_clothes_id"])
//	fmt.Println("============")
//	fmt.Println(c.RegionConfig.ID.ReviewTags.ProductTypeClothesId)
//}
func main() {
	d := `{
"config": {
  "a": "aa",
  "b": "bb",
  "c": "20"
},
"config_test": {
  "a": "aa",
  "b": "bb",
  "c": "20"
}
}`
	// 可以直接将结构体中构建产生一个字段， 然后就只会序列化这一个字段，不会使用更多的字段
	type A struct {
		Endpoint string `json:"a"`
	}
	n := `{
  "a": "aa",
  "b": "bb",
  "c": "20"
}`

	var vd map[string]map[string]string
	err := json.Unmarshal([]byte(d), &vd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vd["config"]["a"])

	var h A
	err = json.Unmarshal([]byte(n), &h)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("===", h.Endpoint)
}
