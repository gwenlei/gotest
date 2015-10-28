package main

import "fmt"

type a interface {
	Sayhi
	Sayhello
}
type Sayhi interface {
	Say()
}
type Sayhello interface {
	Say()
}
type dog struct {
	name string
}
type cat struct {
	name string
}

func (d dog) Say() {
	fmt.Println("hello")
}
func (c cat) Say() {
	fmt.Println("hi")
}
func main() {
	var t a
	var d dog
	var c cat
	fmt.Println(t)
	fmt.Println(d)
	fmt.Println(c)
}
