package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for i := 0; ; i++ {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second / 5)
	}
}
func barnagibiryap(c chan string) {
	for i := 0; ; i++ {
		c <- "1"
	}
}
func barnah(b chan string) {
	m := <-b
	fmt.Println(m)

}
func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
