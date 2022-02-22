package main

import (
	"errors"
	"fmt"
)

type Widget struct {
	X, Y int
}
type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

type Button struct {
	Label // Embedding (delegation)
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//因为这个接口可以通过 Label 的嵌入带到新的结构体，
//所以，可以在 Button 中重载这个接口方法
func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}

func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

type IntSet struct {
	data map[int]bool
	//将undo 做成一个单独的功能，进行反转控制， 而不是直接创建一个新的类型    undoableset 然后override   add  delete   contains 方法
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}
func main() {
	// 如果是embeding 字段就字段的名称就直接是原有的结构体的名字
	label := Label{Widget: Widget{10, 10}, Text: "State:"}

	label.X = 11
	label.Y = 12
	fmt.Println(label.X, label.Widget.X)

	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := Button{Label{Widget{10, 70}, "OK"}}
	//button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}

}
