package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(channel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("*** sending value ***")
	channel <- "a"
	fmt.Println("*** sending value 2 ***")
	channel <- "b"
}

func main() {
	var myChannel chan string
	myChannel = make(chan string)

	go send(myChannel)

	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

}
