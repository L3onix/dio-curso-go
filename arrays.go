package main

import "fmt"

func main(){
  var notas [5]float64
  notas[0] = 5.3
  notas[1] = 8
  notas[2] = 3.4
  notas[3] = 9.2
  notas[4] = 7.9

  var somaNotas float64 = 0
  for i:=0; i < len(notas); i++ {
    somaNotas += notas[i]
  }

  fmt.Println(somaNotas/float64(len(notas)))

  // slices
  fatia := notas[1:4]
  fmt.Println(fatia)

  // append
  fatia = append(fatia, 2.1)
  fmt.Println(fatia)

  // maps
  mapa := make(map[string]int)
  mapa["goleiro"] = 1
  mapa["zagueiro"] = 4
  mapa["lateral"] = 6
  fmt.Println(mapa)
}

