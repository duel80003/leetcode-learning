package main

import (
	"fmt"
)

type iSpeak interface {
	Speak()
}

type dog struct {
	Name string
}

func (d *dog) Speak() {
	fmt.Println("I'm a dog and my name is", d.Name)
}

type cat struct {
	Name string
}

func (c *cat) Speak() {
	fmt.Println("I'm a cat and my name is", c.Name)
}

func factory(t, name string) iSpeak {
	switch t {
	case "dog":
		return &dog{Name: name}
	case "cat":
		return &cat{Name: name}
	default:
		return nil
	}
}

func main() {
	a := factory("cat", "aaa")
	a.Speak()

	b := factory("dog", "ddd")
	b.Speak()
}
