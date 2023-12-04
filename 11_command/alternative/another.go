package alternative

import "fmt"

// https://refactoring.guru/design-patterns/command/go/example#example-0

// [1] Invoker
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// intf
type Command interface {
	execute()
}

// OnCommand - concrete cmd
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// OffCommand - concrete cmd
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// [2] Receiver intf
type Device interface {
	on()
	off()
}

// Concrete receiver
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
