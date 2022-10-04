// this solution for calculating square of perimeter
package main
import ("fmt")

func main() {
	var side_a int
	fmt.Printf("Please enter side of square: ")
	fmt.Scan(&side_a)
  	fmt.Printf("Perimeter of given squeare: %v result is: %v", side_a, (4 * side_a))
}