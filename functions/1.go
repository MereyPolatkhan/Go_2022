package main

import (
	"fmt"
)

func scoreSummary(name string, a, b, c float64) float64 {
	var average float64 = (a + b + c) / 3
	return average

}

func main() {
	var s string
	var a, b, c float64
	fmt.Scan(&s, &a, &b, &c)

	var average = scoreSummary(s, a, b, c)

	fmt.Println("	Name | Score 1 | Score 2 | Score 3 | Average")
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("	%s |	 %0.2f |   %0.2f |   %0.2f | %0.2f", s, a, b, c, average)

}
