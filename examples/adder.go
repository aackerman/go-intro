package main

import "fmt"

func adder(i int) func() int {
  return func() int {
    i += 5
    fmt.Printf("%v\n", i);
    return i
  }
}

func main() {
  addFive := adder(0)
  addFive()
  addFive()
  addFive()
  addFive()
  addFive()
}
