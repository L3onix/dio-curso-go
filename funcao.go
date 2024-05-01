package main

import "fmt"

func media (lista []float64)float64 {
  var total float64
  for _, value := range lista {
    total += float64(value)
  }

  return total / float64(len(lista))
}

func main() {
  notas := []float64{98, 93, 77, 82, 83}

  media := media(notas)
  fmt.Println(media)
}
