package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println("[daqkvz7zec]", err)
		return
	}
	for i := 0; i < 10; i++ {
		index := i
		go func() {
			for {
				time.Sleep(time.Second)
				_, err := w.WriteString("thread [" + strconv.Itoa(index) + "] " + time.Now().String() + "\n")
				if err != nil {
					fmt.Println("[b1zei4zyt9] ", index, err)
				}
			}
		}()
	}
	reader := bufio.NewReader(r)
	n := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic("qi8y2676pt")
		}
		n++
		fmt.Println(n, string(line))
	}
	c := make(chan int, 0)
	<-c
}
