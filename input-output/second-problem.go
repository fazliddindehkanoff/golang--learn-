//this solution for calculating square of area
package main
import ("fmt")

func main(){
	var side_a int
	fmt.Printf("Enter side of square: ")
	fmt.Scan(&side_a)
	fmt.Printf("a side: %v result is: %v", side_a, (side_a*side_a))
}