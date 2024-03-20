// Реализовать паттерн «адаптер» на любом примере.
package main

import (
	"fmt"
)

type Target interface {
	Request() string
}

// адаптируемый класс
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "adaptable method"
}

// адаптер
type Adapter struct {
	adaptee *Adaptee
}

func (ad *Adapter) Request() string {
	return ad.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	result := adapter.Request()
	fmt.Println("result:", result)
}
