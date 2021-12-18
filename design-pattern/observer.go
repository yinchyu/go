package main

import "fmt"

// 收集所有的订阅消息的用户
/*
第一个问题是使用同步阻塞还是异步非阻塞
同步阻塞是最经典的实现方式，主要是为了代码解耦；
异步非阻塞除了能实现代码解耦之外，还能提高代码的执行效率；
第二个问题是如何保证所有Observer都通知成功。
方案一是利用消息队列ACK的能力，Observer订阅消息队列。Subject只需要确保信息通知给消息队列即可。
方案二是Subject将失败的通知记录，方便后面进行重试。
方案三是定义好规范，例如只对网络失败这种错误进行记录，业务失败类型不管理，由业务自行保证成功。
第三个问题是不同进程/系统如何进行通知。
进程间的观察者模式解耦更加彻底，一般是基于消息队列来实现，用来实现不同进程间的被观察者和观察者之间的交互
*/
type subject struct {
	observers []*observer
}

func (su *subject) add(ob *observer) {
	su.observers = append(su.observers, ob)
}

func (su *subject) notify(context string) {
	for _, val := range su.observers {
		// 这个是observers 自己的方法
		val.update(context)
	}
}

func (su *subject) remove(name string) {
	for index, val := range su.observers {
		if val.name == name {
			su.observers = append(su.observers[:index], su.observers[index+1:]...)
			return
		}

	}
}

func newsubject() *subject {
	return &subject{[]*observer{}}

}

type observer struct {
	name string
}

func newobserver(name string) *observer {
	return &observer{name: name}

}

func (ob *observer) update(content string) {
	fmt.Println(ob.name, "receive from ", content)
}

func main() {
	subject := newsubject()
	reader1 := newobserver("reader1")
	reader2 := newobserver("reader2")
	reader3 := newobserver("reader3")
	subject.add(reader1)
	subject.add(reader2)
	subject.add(reader3)
	subject.notify("observer mode")
	subject.remove("reader1")
	subject.remove("reader2")
	subject.remove("reader1")
	subject.notify("observer mode")

}
