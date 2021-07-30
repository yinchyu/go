package main

import "fmt"

type subject struct {
	observers []*observer
}

func  (su *subject) add(ob *observer){
	su.observers=append(su.observers, ob)
}

func  (su *subject) notify(context string){
	for _,val:=range su.observers{
		val.update(context)
	}
}

func  (su *subject) remove(name string){
	for index,val:=range su.observers{
		if val.name==name{
			su.observers=append(su.observers[:index],su.observers[index+1:]...)
			return
		}

	}
}


func newsubject() *subject{
	return &subject{[]*observer{}}

}
type observer struct {
name string
}
func newobserver(name string) *observer{
	return &observer{name: name}

}
func ( ob *observer) update( content string){
	fmt.Println(ob.name,"receive from ",content)
}



func main(){
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

