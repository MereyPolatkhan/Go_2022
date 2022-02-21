package main

import "fmt"

func p(l, w float64) float64 {
	return 2 * (l + w)
}

func main() {
	var l, w float64
	var sum float64 = 0

	for i := 0; i < 3; i++ {
		fmt.Scan(&l, &w)
		sum += p(l, w)
	}

	fmt.Printf("%0.2f", sum)
}
