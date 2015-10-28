package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var c chan int
var c2 chan int
var buf []byte
var tips map[string]string

func main() {
	c = make(chan int)
	c2 = make(chan int)
	buf = make([]byte, 1024)
	tips = map[string]string{
		"name":  "doge",
		"age":   "1",
		"hobby": "programmer",
	}
	inittips()
	fmt.Println("do u want to talk to single dog?")
	go listen()
	<-c
}
func listen() {
	fmt.Printf("%v is listening!\n", tips["name"])
	for {
		go timer()
		os.Stdin.Read(buf)
		if strings.Contains(string(buf), "quit") {
			c <- 1
			break
		}
		go answer()
		<-c2
		buf = make([]byte, 1024)
	}
}
func answer() {
	fmt.Println("i told u!")
	index := strings.IndexRune(string(buf), 0)
	buf = []byte(tips[string(buf)[0:index-1]])
	os.Stdin.Write(buf)
	fmt.Println("\n")
	c2 <- 1
}
func inittips() {
	f, _ := os.Open("/home/code/gotest/tips.txt")
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		onetip := strings.Split(line, ":")
		tips[onetip[0]] = onetip[1]
	}
}
func timer() {
	time.Sleep(time.Second * 5)
	if strings.IndexRune(string(buf), 0) == 0 {
		fmt.Println("i have waited my life time long！")
	}
}
