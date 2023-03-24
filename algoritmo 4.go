package main

import "fmt"

func main() {
var altura float64
var peso float64
var sexo int
fmt.Print("qual sua altura?")
fmt.Scan(&altura)
	fmt.Print("qual seu peso?")
	fmt.Scan(&peso)
	fmt.Print("qual seu sexo?")
	fmt.Scan(&sexo)
imc := peso / (altura * altura)
	switch sexo {
	case 1:
		if imc < 16 {
			fmt.Println(!"magreza grave")
		}
		else if imc < 17.5 {
			fmt.Println(!"magreza grave")
		}
		else if imc < 25 {
			fmt.Println(!"saudavel")
		}
		else if imc < 30 {
			fmt.Println(!"obesidade grau 1")
		}
		else if imc < 35 {
			fmt.Println(!"obesidade grau 2")
		else if imc < 40 {
			fmt.Println(!"obesidade grau 3")
		}
	case 2:

	}
}
