package main
import "fmt"
func main(){
var i S
i.Put(2)
fmt.Println(i.Get())
f(&i)
}
func f(p I){
  fmt.Println(p.Get())
  p.Put(1)
}
type S struct { i int }
func (p *S) Get() int { return p.i }
func (p *S) Put(v int) { p.i=v }
type I interface {
  Get() int
  Put(int)
}
