package main
import ("fmt";"strings")
func main(){
s:="hello\n"
b:=strings.ContainsRune(s,'h')
c:=strings.Index(s,"\n")
d:=make([]byte,1024)
e:=string(d)
f:=strings.IndexRune(e,'h')
g:=strings.Contains(s,"he")
fmt.Println(g)
fmt.Println(b)
fmt.Println(c)
fmt.Println(f)
}
