// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)
package main

import "fmt"

type Human struct {
	Name string
}

func (h *Human) SayHello() {
	fmt.Printf("Hi, my name is %s\n", h.Name)
}

type Action struct {
	Human
}

func (a *Action) Greeting() {
	a.SayHello()
}

func main() {
	h := Human{Name: "Slim Shady"}
	a := Action{Human: h}
	a.Greeting()
}
