// The solution for calculating Gipotenuses and Perimetr of triangle
package main

import (
	"fmt"
	"math"
)

func main() {
	var first_num, second_num, result float64
	fmt.Printf("Enter 2 numbers (seperated by whitespaces): ")
	fmt.Scan(&first_num, &second_num)
	result = math.Sqrt(math.Pow(first_num, 2)+math.Pow(second_num, 2))
	fmt.Printf("Gipotenuses of trinangle is: %v\nPerimeter of triangle is: %v", result, (first_num+second_num+result))
}