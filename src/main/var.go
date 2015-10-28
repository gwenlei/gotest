package main

import (
	"errors"
	"fmt"
)

func main() {
	var v1, v2 = 1, 3
	fmt.Println(v1, v2)
	err := errors.New("emit macho dwar")
	if err != nil {
		fmt.Print(err)
	}
	const (
		x = iota
		y = iota
		z = iota
		w
	)
	fmt.Println(x, y, z, w)
	var array = [3]int{1, 2, 3}
	var slice = []int{0, 3}
	fmt.Println(array, slice)
	slice = array[2:3]
	fmt.Println(slice)
}
