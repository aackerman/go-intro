package main

import "fmt"

func adder(i int) func(int) int {
	return func(k int) int {
		k += i
		fmt.Printf("%v\n", k)
		return i
	}
}

func main() {
	addFive := adder(5)
	addFive(7)
	addFive(4)
}
