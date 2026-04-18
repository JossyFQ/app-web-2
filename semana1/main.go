package main

import "fmt"

func main() {
	// Declaración de variables
	nombre := "Josselyn Fernandez"
	edad := 21
	carrera := "Tecnologías de la Información"
	semestre := 6
	promedio := 8.30

	// Impresión con formato
	fmt.Printf("Soy %s, tengo %d años\n", nombre, edad)
	fmt.Printf("Estudio %s, semestre %d\n", carrera, semestre)
	fmt.Printf("Mi promedio es %.2f\n", promedio)
}