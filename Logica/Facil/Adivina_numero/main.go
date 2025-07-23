package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	contador := 0
	fmt.Println("¡Bienvenido al juego de Adivina el Número!")
	rand.Seed(time.Now().UnixNano()) // Inicializa el generador de números aleatorios
	num_secreto := rand.Intn(100 - 1)+1 // Genera un número aleatorio entre 1 y 100
	for true {
		var numero int
		fmt.Print("Introduce un número entre 1 y 100: ")
		fmt.Scan( &numero)
		contador++ // Incrementa el contador de intentos
		if numero == num_secreto {
			fmt.Printf("¡Felicidades! Has adivinado el número %d en %d intentos.\n", num_secreto, contador)
			break // Sale del bucle si el número es correcto
		}else if numero < num_secreto {
			fmt.Println("El número es mayor. Inténtalo de nuevo.")
		} else {
			fmt.Println("El número es menor. Inténtalo de nuevo.")
		}

	}
}