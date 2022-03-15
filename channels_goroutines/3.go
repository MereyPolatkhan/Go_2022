package main

import "fmt"

func abc(mc chan string) {
	mc <- "a"
	mc <- "b"
	mc <- "c"
}

func def(mc chan string) {
	mc <- "d"
	mc <- "e"
	mc <- "f"
}

func main() {
	myChannel1 := make(chan string)
	myChannel2 := make(chan string)

	go abc(myChannel1)
	go def(myChannel2)

	fmt.Println(<-myChannel1) // 1
	fmt.Println(<-myChannel2) // 2
	fmt.Println(<-myChannel1) // 1
	fmt.Println(<-myChannel2) // 2
	fmt.Println(<-myChannel1) // 1
	fmt.Println(<-myChannel2) // 2

}
