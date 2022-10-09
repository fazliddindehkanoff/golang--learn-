// This solution for calculating circle of area by it's radius
package main
import ("fmt")

const Pi = 3.14

func main() {
	var first_radius, second_radius float64
	fmt.Printf("Enter 2 numbers (seperated by whitespaces): ")
	fmt.Scan(&first_radius, &second_radius)
	fmt.Printf("S1 is equal to: %v\nS2 is equal to: %v\nS1 is equal to: %v", (Pi*first_radius), (Pi*second_radius), (Pi*(first_radius - second_radius)))
}