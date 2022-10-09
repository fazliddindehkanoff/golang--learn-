package main
import (
	"fmt"
	"math"
)

func main(){
	var first_number, second_number float64
	fmt.Printf("Enter two numbers: ")
	fmt.Scan(&first_number, &second_number)
	fmt.Printf("The result is: %v",(math.Abs(first_number-second_number)))
}
