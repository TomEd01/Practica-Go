package main
import (
	"fmt"
)
func main(){
	num := 0
	fmt.Scan(&num)
	suma := 0
	for i:=1; i<=num-1;i++{
		if i%3==0 || i%5==0{
			fmt.Print(i," ")
			suma+=i
		}
	}
	fmt.Println("= ",suma)
}
