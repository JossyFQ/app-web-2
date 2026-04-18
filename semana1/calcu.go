package main

import "fmt"

func main() {

	/*fmt.Printf("==== CALCULADORA CIENTIFICA v1.0 ====")

	var a = 50
	var b = 20

	suma := a + b
	resta := a - b
	multiplicacion := a * b
	divicion := float64(a) / float64(b)

	fmt.Printf("%d + %d = %d\n", a, b, suma)
	fmt.Printf("%d - %d = %d\n", a, b, resta)
	fmt.Printf("%d * %d = %d\n", a, b, multiplicacion)
	fmt.Printf("%d / %d = %.2f\n", a, b, divicion)*/

	//////////////////////////////////////////

	fmt.Println("==== CALCULADORA CIENTIFICA v1.1 ==== \n")

	var numer1 int
	var numer2 int
	//var resultado float64
	var resultado int

	fmt.Printf("Ingresa el primer numero: ")
	fmt.Scanln(&numer1)
	fmt.Printf("\nIngresa el segundo numero: ")
	fmt.Scanln(&numer2) 

	var operacion string
	fmt.Println("\nIngresa la operacion (+, -, *, /, ^, !): ")
	fmt.Scanln(&operacion)
	switch operacion {
		case "+":
			//resultado = float64(numer1) + float64(numer2)
			resultado = numer1 + numer2
		case "-":
			//resultado = float64(numer1) - float64(numer2)
			resultado = numer1 - numer2
		case "*":
			//resultado = float64(numer1) * float64(numer2)
			resultado = numer1 * numer2
		case "/":
			if numer2 != 0 && numer1 != 0 {
				resultado = numer1 / numer2
				//resultado = float64(numer1) / float64(numer2)
			}else{
				fmt.Println("\n'Error: no se puede dividir entre cero'")
			}
		/////////////////////////////
		case "^":
			if numer2 == 0{
				resultado = 1
			}else{
				resultado = 1
				for i := 0; i < int(numer2); i++{
					resultado *= numer1
				}
			}
		case "!":
			if numer1 == 0{
				resultado = 1
			}else{
				resultado = 1
				for i := 1; i <= int(numer1); i++{
					//resultado *= float64(i)
					resultado *= i
				}
			}
	}
	fmt.Printf("\nResultado: %d", resultado)

	//////////////////////////////////////////
	
}