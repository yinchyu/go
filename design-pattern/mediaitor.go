//package  main
//
//import "fmt"
//
//type Sound struct {
//	common Driver
//}
//
//type Video struct {
//	common Driver
//}
//
//func (S *Sound)Show(){
//	S.common.Process(S)
//}
//func (V *Video) Show()  {
//	V.common.Process(V)
//}
//
//
//type Driver struct {
//	CD  *Sound
//	VIDEO *Video
//}
//
//func (d *Driver)Process(i interface{}){
//	switch i.(type) {
//	case *Sound:
//		fmfmt.Println("音乐播放器")
//	case *Video:
//		fmfmt.Println("视频播放器")
//	default:
//		fmfmt.Println("没有确定的类型")
//	}
//}
//
//
//func main()  {
//	d:=Driver{}
//	s:=Sound{d}
//	v:=Video{d}
//	d.CD=&s
//	d.VIDEO=&v
//	s.Show()
//	v.Show()
//
//
//
//
//
//}
package main

import (
	"fmt"
	"strings"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData() {
	c.Data = "music,image"

	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	GetMediatorInstance().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediatorInstance().changed(s)
}

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	default:
		fmt.Println("=========",i)
	}

}

func main() {
	mediator := GetMediatorInstance()
	mediator.CD = &CDDriver{}
	mediator.CPU = &CPU{}
	mediator.Video = &VideoCard{}
	mediator.Sound = &SoundCard{}

	//Tiggle
	mediator.CD.ReadData()
	//
	//if mediator.CD.Data != "music,image" {
	//	fmt.Printf("CD unexpect data %s", mediator.CD.Data)
	//}
	//
	//if mediator.CPU.Sound != "music" {
	//	fmt.Printf("CPU unexpect sound data %s", mediator.CPU.Sound)
	//}
	//
	//if mediator.CPU.Video != "image" {
	//	fmt.Printf("CPU unexpect video data %s", mediator.CPU.Video)
	//}
	//
	//if mediator.Video.Data != "image" {
	//	fmt.Printf("VidoeCard unexpect data %s", mediator.Video.Data)
	//}
	//
	//if mediator.Sound.Data != "music" {
	//	fmt.Printf("SoundCard unexpect data %s", mediator.Sound.Data)
	//}
}
