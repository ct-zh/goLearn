package main

import "fmt"

type Message struct {
	msg string
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{msg: msg}
}

type Greeter struct {
	Message Message
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func main() {
	//
	//message := NewMessage("hello world")
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//event.Start()

	event := InitializeEvent("hello world")
	event.Start()
}
