package main

import "fmt"

const n string = "Este valor no se cambia" //Tambien se puede defiinir afuera de la funcion main
func main() {
	//Variables
	var nombre string //String = cadena de caracteres
	nombre = "Esto es una cadena de caracteres" //Se puede modificar a lo largo del codigo
	fmt.Println(nombre)//String
	fmt.Println("\n")
	/*----------------------------------------------------*/
	var edad, edad2 int = 19, 3//Int = entero
	fmt.Println(edad,edad2,"Soy dos enteros")//int
	fmt.Println("\n")
	/*----------------------------------------------------*/
	flotante := 3.6 //Float = tambien se puede definir así pero cuando estamos seguros
	flotante = 5.5 //Tambien
	fmt.Println(flotante,"Soy un flotante")//float
	/*----------------------------------------------------*/
	fmt.Println("\n")
	var (//Asi se define muchas variables
		pi float32
		bo  bool
		cadena = "Texto 2"
		fecha = 12032022
	)
	pi =3.1416
	bo= true
	fmt.Println(pi,bo,cadena,fecha)
	fmt.Println("\n","Soy una constante y ",n)
}