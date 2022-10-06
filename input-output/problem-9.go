//This solution for calculating geometrics of 2 numbers
package main
import (
	"fmt"
	"math"
)

func main() {
	var first_num, second_num float64
	fmt.Printf("Enter two number (seperated by whitespaces): ")
	fmt.Scan(&first_num, &second_num)
	fmt.Printf("Result is: %v", (math.Sqrt((first_num*second_num))))
}