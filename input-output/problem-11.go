package main

import (
	"fmt"
	"math"
)

func main() {
	var first_num, second_num float64
	fmt.Printf("Enter 2 numbers (seperated by whitespaces): ")
	fmt.Scan(&first_num, &second_num)
	fmt.Printf("Results: addition of nums: %v\nMultiplication of them: %v\nModule of %v = %v\nModule of %v = %v", (first_num + second_num),(first_num * second_num), first_num, math.Abs(first_num), second_num, math.Abs(second_num))
}