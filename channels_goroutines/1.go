package main

import (
	"errors"
	"fmt"
	"log"
)

func calc(r int) (int, error) {
	if r < 0 {
		return 0, errors.New("provide positive number")
	}
	return r * r, nil
}

func sayHello() {
	fmt.Println("Hello________++++++++++++++++_______++++++++++++++")
}
func main() {

	defer sayHello()

	area, err := calc(-9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(area)

}

// why it is printed BEFORE the PANIC
