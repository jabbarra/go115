package main

import (
	"fmt"
	"time"
)

func executeChannelsExample() {

	channel := make(chan string)

	go echoPingServer(channel)
	go sendPingClient(channel)

	var input string
	fmt.Scanln(&input)
	fmt.Println("End  example")
}

func echoPingServer(channel chan string) {
	var counter int
	for {
		counter++
		fmt.Println(<-channel, " ", counter)
		time.Sleep(time.Second * 1)
	}
}

func sendPingClient(channel chan string) {
	for {
		channel <- "ping"
	}
}