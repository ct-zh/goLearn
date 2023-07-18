package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Message string

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message(msg)
}

type Greeter struct {
	Message Message
	Grumpy  bool
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
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
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func main() {
	//
	//message := NewMessage("hello world")
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//event.Start()

	//event := InitializeEvent("hello world")
	//event.Start()

	e, err := InitializeEvent("hello world")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
