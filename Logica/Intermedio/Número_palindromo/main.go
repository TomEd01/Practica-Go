package main
import (
	"fmt"
)
func main() {
	num := 0
	fmt.Scan(&num)
	original := num // Guardar el número original
	reversed := 0    // Variable para almacenar el número invertido

	for num > 0 {
		digit := num % 10 // Obtener el último dígito
		reversed = reversed*10 + digit // Construir el número invertido
		num = num / 10 // Eliminar el último dígito del número
	}

	if original == reversed {
		fmt.Println("Es un palíndromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}