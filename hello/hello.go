package main

import (
	"fmt"
	"time"
)

func main() {
	var a, b = "Jakob", "H"
	fmt.Printf("hello, world\n")
	fmt.Println("Time:", time.Now())
	fmt.Println("Result:", sub(10, 5))
	if read := false; read {
		for i := 0; i < 5; i++ {
			fmt.Println("Strings: ", a, b)
			a, b = swap(a, b)
			fmt.Println("Swapped: ", a, b)
		}
	} else {
		fmt.Println("Else invoked.")
	}
}

func sub(x, y int) int {
	return x - y
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}
