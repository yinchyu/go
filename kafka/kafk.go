package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-ini/ini"
	"github.com/hpcloud/tail"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func sendmessage(){
	fmt.Printf("start\n")
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	// 构造⼀个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"119.45.176.15:9092"},
		config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

func readtailfile(){
	fileName := "./my.log"
	config:=tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		msg *tail.Line
		ok bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n",
				tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg: ", msg.Text)
		fmt.Println("time: ", msg.Time)
	}
}
func readconfig(){
	type Embeded struct {
		Dates  []time.Time `delim:"|" comment:"Time data"`
		Places []string    `ini:"places,omitempty"`
		None   []int       `ini:",omitempty"`
	}

	type Author struct {
		Name      string `ini:"NAME"`
		Male      bool
		Age       int `comment:"Author's age"`
		GPA       float64
		NeverMind string `ini:"-"`
		*Embeded `comment:"Embeded section"`
	}


	a := &Author{"Unknwon", true, 21, 2.8, "",
		&Embeded{
			[]time.Time{time.Now(), time.Now()},
			[]string{"HangZhou", "Boston"},
			[]int{},
		}}
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, a)
	if err!=nil{
		fmt.Println(err)
	}
	cfg.SaveTo("./test.ini")
}
func writetsddata(){
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "dYUjfehh3g4yHmcCoO5pw8bbGCdZwUq7_FOsA4HtzMb0-JHJCDUeE-An-tosS6L8Lsj5DMEo2nNwTO4vpo2tlA=="
	const bucket = "test"
	const org = "ycy"

	client := influxdb2.NewClient("http://www.yinchangyu.top:8086", token)
	// always close client at the end
	defer client.Close()
	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)
    //  获取对应的cpu 信息

	// a:=fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	// b:=fmt.Sprintf("stat,Percent=%f",c2[0])
	// fmt.Println(a,b)
	// write line protocol
	for{
		//  写整个数据， 可以写一个点数据
		// writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
		c2, _ := cpu.Percent(time.Duration(time.Second), false)
		c3,_:=mem.VirtualMemory()
		c4,_:=net.IOCounters(true)


		// fmt.Println(c2)
		// point:=write.Point{
		// 	"packagetest",
		// 	[]*protocol.Tag
		//
		// }
		// writeAPI.WritePoint(&point)
		//  写入不同行的数据
		writeAPI.WriteRecord(fmt.Sprintf("cpu avg=%f",c2[0]))
		writeAPI.WriteRecord(fmt.Sprintf("memery total=%d,available=%d,used=%d",c3.Total,c3.Available,c3.Used))
		writeAPI.WriteRecord(fmt.Sprintf("net send=%d,recv=%d",c4[0].BytesSent,c4[0].BytesRecv))
		// Flush writes
		// time.Sleep(time.Millisecond*100)
		fmt.Println("write data",time.Now())
		writeAPI.Flush()
	}

}

func main(){
	// sendmessage()
	// readconfig()
	writetsddata()
}


