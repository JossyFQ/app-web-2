package main

import "fmt"

func main() {

	fmt.Printf("==== CALCULADORA CIENTIFICA v1.0 ==== \n")

	var a = 50
	var b = 20

	suma := a + b
	resta := a - b
	multiplicacion := a * b
	divicion := float64(a) / float64(b)

	fmt.Printf("%d + %d = %d\n", a, b, suma)
	fmt.Printf("%d - %d = %d\n", a, b, resta)
	fmt.Printf("%d * %d = %d\n", a, b, multiplicacion)
	fmt.Printf("%d / %d = %.2f\n", a, b, divicion)

}