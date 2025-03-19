package main

import "fmt"

const TF float32 = 32

func main() {
	var TC float32 = TF - 32
	var TK float32 = TF + 241
	var resp string

	fmt.Println("Digite a medida termometrica desejada: (F/K/C) ")
	fmt.Scanln(&resp)

	if resp == "C" {
		var Ebulicao = TC + 100
		fmt.Printf("A temperatura base em Celsius é de %f, e a temperatura de ebulição é de %f", TC, Ebulicao)
	} else if resp == "K" {
		var Ebulicao = TK + 100
		fmt.Printf("A temperatura base em Kelvin é de %f, e a temperatura de ebulição é de %f", TK, Ebulicao)
	} else if resp == "F" {
		var Ebulicao = TF + 100
		fmt.Printf("A temperatura base em Firehight é de %f, e a temperatura de ebulição é de %f", TF, Ebulicao)
	} else {
		fmt.Println("Resposta invalida.")
	}
}
