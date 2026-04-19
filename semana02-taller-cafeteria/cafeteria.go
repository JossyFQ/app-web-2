package main

import "fmt"

type Cliente struct{
	ID int
	Nombre string
	Carrera string
	Saldo float64
}

type Producto struct{
	ID int
	Nombre string
	Precio float64
	Stock int
	Categoria string
}

type Pedido struct{
	ID int
	ClienteID int
	ProductoID int
	Cantidad int
	Total float64
	Fecha string
}

func ListaClientes(cliente []Cliente){
	fmt.Println("\n====Registro de clientes=====")
	if len(cliente) == 0 {
		fmt.Println("(no hay cliente)")
		return
	}
	for _, c := range cliente {
		fmt.Printf("ID: %d | Nombre: %s | Carrera: %s | Saldo: %.2f\n", c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}
func ListaProducto(producto []Producto) {
		fmt.Println("\n====Registro de Productos=====")
	if len(producto) == 0 {
		fmt.Println("(no hay producto)")
		return
	}
	for _, p := range producto {
		fmt.Printf("ID: %d | Nombre: %s | Precio: %.2f | Stock: %d | Categoria: %s\n", p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
}

func AgregarCliente(cliente []Cliente, nuevo Cliente) []Cliente {
	return append(cliente, nuevo)
}
func AgregarProducto(producto []Producto, nuevo Producto) []Producto {
	return append(producto, nuevo)
}

func BuscarClientePorID(cliente []Cliente, ID int) int {
		for i, c := range cliente {
		if c.ID == ID {
			return i
		}
	}
	return -1
}
func BuscarProducto(producto []Producto, ID int) int {
		for i, p := range producto {
		if p.ID == ID {
			return i
		}
	}
	return -1
}

func EliminarCliente(cliente []Cliente, ID int ) []Cliente {
	idx := BuscarClientePorID(cliente, ID)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", ID)
		return cliente
	}
	return append(cliente[:idx], cliente[idx+1:]...)
}
func EliminarProducto(producto []Producto, ID int ) []Producto {
	idx := BuscarProducto(producto, ID)
	if idx == -1 {
		fmt.Printf("⚠ Producto con ID %d no existe.\n", ID)
		return producto
	}
	return append(producto[:idx], producto[idx+1:]...)
}

func main() {
	cliente := []Cliente{
		{ID: 1, Nombre: "Pedro Sanches", Carrera: "TI", Saldo: 10.00},
		{ID: 2, Nombre: "Luis Moreira", Carrera: "Software", Saldo: 20.00},
		{ID: 3, Nombre: "Maria Cabeza", Carrera: "Enfermeria", Saldo: 40.00},
	}
	producto := []Producto{
		{ID: 1, Nombre: "Pan", Precio: 1.30, Stock: 2, Categoria: "snack"},
		{ID: 2, Nombre: "Batido de mora", Precio: 1.50, Stock: 3, Categoria: "bebida"},
		{ID: 3, Nombre: "Pollo al horno", Precio: 2.50, Stock: 5, Categoria: "almuerzo"},
		{ID: 4, Nombre: "hamburguesa", Precio: 1.15, Stock: 10, Categoria: "snack"},
	}
	//var pedido []Pedido

	ListaClientes(cliente)

	idx := BuscarClientePorID(cliente, 3)
	if idx == -1 {
		fmt.Println ("No existe")
		return
	}
	fmt.Println ("Hola", cliente[idx].Nombre)

	nuevoCliente := Cliente{4, "Byron Casteñeda", "Abogado", 45.00}
	cliente = AgregarCliente(cliente, nuevoCliente)

	EliminarCliente(cliente, 1)

	ListaProducto(producto)

	idx = BuscarProducto(producto, 4)
	if idx == -1 {
		fmt.Println("No existe")
		return
	}	
	fmt.Println("Producto: ", producto[idx].Nombre)

	nuevoProducto := Producto{5, "Tostada", 1.00, 45, "Almuerzo"}
	producto = AgregarProducto(producto, nuevoProducto)

	EliminarProducto(producto, 1)
}
