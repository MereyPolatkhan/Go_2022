package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Print("b")
	}
}

func c() {
	for i := 0; i < 10; i++ {
		fmt.Print("c")
	}
}

func main() {
	go a()
	go b()
	go c()

	time.Sleep(time.Second)
	fmt.Println("\nend main()")

}
