package main

import (
	"fmt"
	"time"
)

func a(channel chan string) {
	for i := 0; i < 10; i++ {
		channel <- "a"
	}
}

func b(channel chan string) {
	for i := 0; i < 10; i++ {
		channel <- "b"
	}
}

func main() {
	myChannel := make(chan string)
	go a(myChannel)
	go b(myChannel)
	time.Sleep(3 * time.Second)

	for i := 0; i < 20; i++ {
		fmt.Print(<-myChannel, " ")
	}

	fmt.Println("\nend main()")
}
