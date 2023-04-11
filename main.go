/*
author: Khaled Alam <khaledalam.net@gmail.com>
created: April-2023

description: Simple ping pong using goroutine.
*/
package main

import (
	"fmt"
	"time"
)

func pingFun(ping <-chan int, pong chan<- int) {
	for {
		<-ping
		fmt.Println("Ping")
		time.Sleep(1 * time.Second)
		pong <- 1
	}
}

func pongFun(ping chan<- int, pong <-chan int) {
	for {
		<-pong
		fmt.Println("Pong")
		time.Sleep(1 * time.Second)
		ping <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pingFun(ping, pong)
	go pongFun(ping, pong)

	ping <- 1

	select {}
}
