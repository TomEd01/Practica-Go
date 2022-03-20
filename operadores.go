package main

import "fmt"

func main() {
	var a, b int
	//Entrada
	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&a)
	fmt.Print("Ingrese otro numero: ")
	fmt.Scanln(&b)

	//Proceso
	fmt.Println("Suma: ", a+b)
	fmt.Println("Resta: ", a-b)
	fmt.Println("División: ", a/b)
	fmt.Println("Multiplicación: ", a*b)
	fmt.Println("Modulo: ", a%b)

	//Operadores
	fmt.Println("¿Son iguales? => ", a == b)
	fmt.Println("¿Son Distintos? => ", a != b)
	fmt.Println("¿El primero es menor que el segundo? => ", a < b)
	fmt.Println("¿El primero es menor o igual que el segundo? => ", a <= b)
	fmt.Println("¿El primero es mayor que el segundo? => ", a > b)
	fmt.Println("¿El primero es mayor o igual que el segundo? => ", a >= b)
}
