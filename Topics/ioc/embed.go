package ioc

import "fmt"

// Painter 把组件画出来
type Painter interface {
	Paint()
}

// Clicker 用于表明点击事件
type Clicker interface {
	Click()
}

type Widget struct {
	X, Y int
}
type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

func (label Label) Paint() {
	// TODO: pointer
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

// TODO: 如果 Button.Paint() 不实现的话，会调用 Label.Paint() ，
// 所以，在 Button 中声明 Paint() 方法，相当于 Override

type Button struct {
	Label // Embedding (delegation)
}

func NewButton(x, y int, text string) Button {
	// TODO: if return is struct type, then should not return point type
	return Button{
		Label: Label{
			Widget: Widget{
				X: x,
				Y: y,
			},
			Text: text,
		},
	}
}

// TODO: 因为这个接口可以通过 Label 的嵌入带到新的结构体，
// 所以，可以在 Button 中重载这个接口方法

func (button Button) Paint() { // TODO: Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}

func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}

// ---

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}

func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

func RunDemo() {
	label := Label{Widget{10, 10}, "State:"}

	label.X = 11
	label.Y = 12

	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	fmt.Println("-------")

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		// everyone can print
		widget.(Painter).Paint()
		// not everyone can click
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
}
