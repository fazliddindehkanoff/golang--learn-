//This solution for calculating average of 2 numbers
package main
import ("fmt")

func main() {
	var first_number, second_number float32
	fmt.Printf("Enter 2 numbers (seperated by whitespaces): ")
	fmt.Scan(&first_number, &second_number)
	fmt.Printf("Result is: %v", (first_number + second_number)/2) 
}