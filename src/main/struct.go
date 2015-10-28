package main

import "fmt"

func main() {
	P := new(person)
	P.name = "jim"
	P.age = 10
	fmt.Println(P)
	S := new(Student)
	S.person.name = "tom"
	S.speciality = "sleep"
	S = &Student{*P, "play", Skills{"c", "d"}}
	fmt.Println(S)
	Sk := Skills{"a", "b"}
	Sk = append(Sk, "hoo")
	fmt.Println(Sk)
}

type person struct {
	name string
	age  int
}
type Student struct {
	person
	speciality string
	Skills
}
type Skills []string
