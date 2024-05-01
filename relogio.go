package main

import (
  "fmt"
  "time"
)

func main() {
  for hora:=0; hora<1; hora++{
    for minuto:=0; minuto<60; minuto++ {
      for segundo:=0; segundo<60; segundo++{
        time.Sleep(time.Second)
        fmt.Print("H:", hora, ":M:", minuto, ":S:", segundo, "\n")
      }
    }
  }
}
