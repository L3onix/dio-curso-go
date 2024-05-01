package main

import "fmt"

func main() {
  for i := 0; i < 5; i++ {
    switch i {
      case 0: fmt.Println("primeiro")
      case 2: fmt.Println("dois")
    }
  }
}
