package main

import (
	"fmt"
	"math"
)

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", r2.area())
	fmt.Println("Area of c1 is:", c1.area())
	fmt.Println("Area of c2 is:", c2.area())
	c1.setrad(1)
	fmt.Println("radius of c1 is:", c1.radius)
	c1.setrad2(2)
	fmt.Println("radius of c1 is:", c1.radius)
	bc := bigCircle{Circle{5}, 6}
	fmt.Println(bc)
	bc.setrad(3)
	fmt.Println(bc)
	fmt.Println("Area of bc is:", bc.area())
	bc.Sayhi()

}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}
type bigCircle struct {
	Circle
	bigrad float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c *Circle) setrad(r float64) {
	c.radius = r
}
func (c Circle) setrad2(r float64) {
	c.radius = r
}
func (c Circle) Sayhi() {
	fmt.Println("hello")
}

//func (c bigCircle) Sayhi() {
//	fmt.Println("hello,bc")
//}
