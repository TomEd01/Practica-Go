package main

import "fmt"

func main() {
	var a int
	fmt.Print("Ingresa un numero: ")
	fmt.Scanln(&a)
	if a==0{
		fmt.Print("Es neutro :p")
	}else if a%2==0{
		fmt.Print("Es par xd")
	}else{
		fmt.Print(a," es impar :(")
	} 
}
