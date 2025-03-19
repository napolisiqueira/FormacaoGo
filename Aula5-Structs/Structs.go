package main 

import "fmt"

type retangulo struct {
  base, altura int
}

func (r *retangulo) area() int {
  return r.base * r.altura
}

func (r retangulo) perimetro() int {
  return (r.base * 2) + (r.altura * 2) 
}

func main () {
  var x, y int
  
  fmt.Println("Digite o valor da base e da altura: ")
  fmt.Scanln(&x, &y)
  r := retangulo{base: x, altura: y}

  fmt.Printf("Area: %d\n",r.area())
  fmt.Printf("Perimetro: %d\n", r.perimetro())
}








