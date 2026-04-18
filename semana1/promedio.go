package main

import "fmt"

func main() {
	var nombre string
	var nota1, nota2 float64

	// Pedir datos al usuario
	fmt.Print("¿Cómo te llamas? ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingresa tu nota 1: ")
	fmt.Scanln(&nota1)

	fmt.Print("Ingresa tu nota 2: ")
	fmt.Scanln(&nota2)

	// Calcular promedio
	promedio := (nota1 + nota2) / 2

	// Mostrar resultados
	fmt.Printf("%s, tu promedio es: %.2f\n", nombre, promedio)

	// Evaluar estado
	if promedio >= 7 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}
}