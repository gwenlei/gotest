package main

import "fmt"

func main() {
	m := months{"January": 31, "February": 28, "December": 31}
	fmt.Println(m)
	for i, _ := range m {
		fmt.Println(m[i])
	}
}

type months map[string]int
