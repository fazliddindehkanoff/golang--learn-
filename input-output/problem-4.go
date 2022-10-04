//this solution for calculating circle of length
package main
import ("fmt")

const Pi = 3.14

func main() {
	var diameter float64
	fmt.Printf("Entere diameter of circle: ")
	fmt.Scan(&diameter)
	fmt.Printf("The result is: %v for diameter: %v", (Pi * diameter), (diameter))
}