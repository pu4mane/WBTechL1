// Разработать программу,
// которая в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}
package main

import (
	"fmt"
	"reflect"
)

func getType(variable interface{}) string {
	t := reflect.TypeOf(variable)

	switch t.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	arr := []interface{}{1, "boom", true, make(chan int, 0), 13.6}
	for _, elem := range arr {
		fmt.Println("elem:", elem, "\ttype:", getType(elem))
	}
}
