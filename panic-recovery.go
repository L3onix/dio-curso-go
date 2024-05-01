// panic: utilizado para report de erros
// recover: captura o trigger de panic

package main

import "fmt"

func main() {
  defer func () {
    x := recover()
    fmt.Println(x)
  } ()
  panic("paniccc")
}
