package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func func1() {
	defer wg.Done()

	fmt.Println("1st goroutine sleeping...")
	time.Sleep(1 * time.Second)
}

func func2() {
	defer wg.Done()

	fmt.Println("2nd goroutine sleeping...")
	time.Sleep(1 * time.Second)
}

func func3() {
	defer wg.Done()

	fmt.Println("3th goroutine sleeping...")
	time.Sleep(1 * time.Second)
}

func func4() {
	defer wg.Done()

	fmt.Println("4th goroutine sleeping...")
	time.Sleep(1 * time.Second)
}

func main() {

	wg.Add(1)
	func1()

	wg.Add(1)
	func2()

	wg.Add(1)
	func3()

	wg.Add(1)
	func4()

	fmt.Println("All gorountines complete.")

	wg.Wait()

}
