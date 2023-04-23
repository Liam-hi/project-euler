package main

import "fmt"

func main() {
	var multiples int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples += i
		}
	}
	fmt.Println(multiples)
}
