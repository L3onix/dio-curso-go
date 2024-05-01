package main

import "fmt"

// declarada fora das funções
type Pessoa struct {
  nome string
  idade int
}

type Geometrico interface {
  getArea() float64
}

type Triangulo struct {
  base float64
  altura float64
}

func (t Triangulo)getArea() float64 {
  // (base * altura) / 2
  return (t.base * t.altura) / float64(2)
}

func printArea(g Geometrico) {
  fmt.Println(g.getArea())
}

func main() {
  var me Pessoa
  me.nome = "Leonardo"
  me.idade = 28
  fmt.Println(me)

  var elemento Triangulo = Triangulo{1.5, 2.3}
  printArea(elemento)
}
