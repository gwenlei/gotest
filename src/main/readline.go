package main
import ("fmt";"os";"strings";"bufio";"io")
func main(){
f,_ :=os.Open("/etc/passwd")
defer f.Close()
buf:=bufio.NewReader(f)
for {
      line,err := buf.ReadString('\n')
      line=strings.TrimSpace(line)
      if err== io.EOF{
        break
      }
  os.Stdout.Write([]byte(line))
  fmt.Println()
      }

}
