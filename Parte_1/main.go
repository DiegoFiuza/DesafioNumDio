package main

import "fmt"

var i int

func main() {
	for i = 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
