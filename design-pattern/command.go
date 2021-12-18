package main

import "fmt"

type Command interface {
	Execute()
}

type Boxer interface {
	pressbutton1()
	pressbutton2()
}

type box struct {
	Command1 Command
	Command2 Command
}

func newBox() *box {
	return &box{}
}

type startcommand struct {
}

func (s startcommand) Execute() {
	fmt.Print("system starting\n")
}

type rebootcommand struct {
}

func (s rebootcommand) Execute() {
	fmt.Print("system rebooting\n")
}
func (b box) pressbutton1() {
	b.Command1.Execute()
}

func (b box) pressbutton2() {
	b.Command2.Execute()
}

func main() {

	start := startcommand{}
	reboot := rebootcommand{}
	motherboard := newBox()
	motherboard.Command1 = start
	motherboard.Command2 = reboot
	motherboard.pressbutton1()
	motherboard.pressbutton2()

	motherboard2 := newBox()
	motherboard2.Command1 = reboot
	motherboard2.Command2 = start
	motherboard2.pressbutton1()
	motherboard2.pressbutton2()

}
