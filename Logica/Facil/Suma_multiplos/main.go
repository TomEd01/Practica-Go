package main
import (
	"fmt"
)
func main(){
	num := 0
	fmt.Scan(&num)
	suma := 0
	for i:=1; i<=num-1;i++{ // Cambiado de num a num-1 para sumar hasta el número anterior
		if i%3==0 || i%5==0{ // Verifica si es múltiplo de 3 o 5
			fmt.Print(i," ")
			suma+=i // Suma el número si es múltiplo de 3 o 5
		}
	}
	fmt.Println("= ",suma)
}
