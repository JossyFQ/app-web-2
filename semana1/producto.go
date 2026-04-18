package main

import "fmt"

func main() {
	// Declaración de constantes
	const producto1 = 49.99
	const producto2 = 29.99
	const producto3 = 15.50

	// Cálculos
	total := producto1 + producto2 + producto3
	promedio := total / 3
	totalConDescuento := total * 0.85

	// Impresión con formato
	fmt.Printf("Producto 1: $%.2f\n", producto1)
	fmt.Printf("Producto 2: $%.2f\n", producto2)
	fmt.Printf("Producto 3: $%.2f\n", producto3)
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Printf("Promedio: $%.2f\n", promedio)
	fmt.Printf("Total con 15%% descuento: $%.2f\n", totalConDescuento)
}