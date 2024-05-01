package main

import "fmt"

func setZero(xPtr *int) {
  *xPtr = 0
}

func main() {
  x := 5
  setZero(&x)
  fmt.Println(x)
}
