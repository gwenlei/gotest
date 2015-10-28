package main

import (
	"fmt"
	"reflect"
	//	"strconv"
)

type empty interface{}
type List []empty

func main() {
	a := 1
	b := "hehe"
	emptyfunc(a)
	emptyfunc(b)
	var c interface{}
	var i int = 5
	s := "hello world"
	c = i
	fmt.Println(c)
	if value, ok := c.(int); ok {
		fmt.Printf("int is %d\n", value)
	}
	fmt.Println(reflect.TypeOf(c))
	c = s
	fmt.Println(c)
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is an int and value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of different type", index)
		}
	}
}
func emptyfunc(e empty) {
	fmt.Println(e)
}
