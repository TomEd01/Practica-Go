package main

import "fmt"

func main() {
	var (
		nombre string
		edad   int
		pi     float64
	)
	//MOSTRAR DATOS POR PANTALLA
	//fmt.Print("Hola \tMundo\n") //Respeta el salto de linea
	//fmt.Println("Hola")           //No respeta el salto de linea
	//fmt.Println("Nombre: ",nombre, "Edad",edad, "\nPi",pi) //Formato bebe
	fmt.Print("Ingrese su Nombre: ")
	fmt.Scanln(&nombre) //Uso stricto

	fmt.Print("Ingrese su Edad: ")
	fmt.Scanln(&edad) //Uso flexible

	fmt.Print("Ingrese su el valor de Pi: ")
	fmt.Scanln(&pi) //Uso stricto
	/*
		fmt.Printf("Ingresa nombre: ")
		fmt.Scanf("%s", &nombre)

		fmt.Printf("Ingresa edad: ")
		fmt.Scanf("%d", &edad)

		fmt.Printf("Ingresa Pi: ")
		fmt.Scanf("%f", &pi)
	*/
	fmt.Printf("Nombre: %s \t Edad: %d \nPi: %f", nombre, edad, pi) //Formato C = formatear los datos
}
