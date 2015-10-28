package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int
}

func main() {
	var p Person = Person{"tom", 12}
	t := reflect.TypeOf(&p)
	v := reflect.ValueOf(&p)
	fmt.Println(t, v)
	fmt.Println(t.Elem(), v.Elem())
	tag := t.Elem().Field(0).Tag
	name := v.Elem().Field(0).String()
	fmt.Println(tag, name)
	tag = t.Elem().Field(1).Tag
	age := v.Elem().Field(1).Int()
	fmt.Println(tag, age)
}
