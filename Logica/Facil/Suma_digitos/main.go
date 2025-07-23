package main
import (
	"fmt"
)
func main(){
	num := 0
	fmt.Scan(&num)
	suma := 0
	for i:=0; num>0;i++{
		ultimo_digito := 0
		ultimo_digito = num%10
		suma+=ultimo_digito
		num = num /10
	}
	fmt.Print(suma)
}