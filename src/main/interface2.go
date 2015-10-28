package main
import ("fmt";"strconv")
type Human struct{
name string
age int
phone string
}
func (h Human) String() string {
return "<"+h.name+"-"+strconv.Itoa(h.age)+"years - "+h.phone+">"
}
func main(){
Bob:=Human{"Bob",39,"000-777-XXX"}
fmt.Println("this human is : ",Bob)
fmt.Println("this human is : ",2)
var a interface{}
var i int=5
s:="hello world"
a=i
a=s
fmt.Println(a)
}
