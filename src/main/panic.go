package main

import "fmt"

func main() {
	fmt.Println(throwsPanic(calm))
}
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}
func calm() {
	panic("try panic")
	fmt.Println("calm down please")
}
