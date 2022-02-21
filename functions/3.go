package main

import (
	"errors"
	"fmt"
)

func p2(l, w float64) (float64, error) {
	if l < 0 {
		var s1 string = fmt.Sprintf("a length of %f is invalid", l)
		var err1 error = errors.New(s1)
		return 0, err1
	}
	if w < 0 {
		var s2 string = fmt.Sprintf("a width of %0.2f is invalid", w)
		var err2 error = errors.New(s2)
		return 0, err2
	}
	return 2 * (l + w), nil
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	p, err := p2(a, b)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%0.2f", p)
	}
}
