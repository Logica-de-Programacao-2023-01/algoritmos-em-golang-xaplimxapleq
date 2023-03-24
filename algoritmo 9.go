package main

import "fmt"

func main() {
	var soma, quantidade, x int
	fmt.Print("digite o valor do numero")
	fmt.Scan(&x)
	soma += x
	quantidade++
	for x != 0 {
		fmt.Print("digite um numero")
		fmt.Scan(&x)
		soma += x
		if x != 0 {
			soma += x
			quantidade++
			media := soma / quantidade
			fmt.Println("a média é de", media)
		}
	}

}
