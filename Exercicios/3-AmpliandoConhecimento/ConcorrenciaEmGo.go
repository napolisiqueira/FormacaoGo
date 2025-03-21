package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "Ping"
	}
}

func pong(c chan string) {
	for {
		c <- "Pong"
	}
}

func print(c chan string) {
	for {
		message := <-c
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go print(c)

	fmt.Scanln()

}
