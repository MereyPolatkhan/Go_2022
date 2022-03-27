package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	work := func(id int) {
// 		defer wg.Done()
// 		fmt.Printf("Goroutines %d start work\n", id)
// 		time.Sleep(2 * time.Second)
// 		fmt.Printf("Goroutines %d end work\n", id)
// 	}

// 	go work(1)
// 	go work(2)
// 	// go work(3)

// 	wg.Wait()

// 	fmt.Println("Goroutines done work")
// }
