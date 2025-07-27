package main
import (
	"fmt"
)
func main(){
	num := 0
	fmt.Scan(&num)
	suma := 0
	for i:=0; num>0;i++{
		ultimo_digito := 0 // Variable para almacenar el último dígito
		ultimo_digito = num%10 // Obtener el último dígito
		suma+=ultimo_digito // Sumar el último dígito a la suma total
		num = num /10 // Eliminar el último dígito del número
	}
	fmt.Print(suma)
}