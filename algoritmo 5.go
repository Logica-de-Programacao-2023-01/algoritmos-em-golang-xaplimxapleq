package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("digite o numero")
	fmt.Scan(&x)
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("é multiplo de 3 e 5")
	} else {

		fmt.Println("não é multiplo de 5")
	}

}
