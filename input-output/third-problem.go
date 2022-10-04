// this solution for calculating rectangle of area and its perimeter
package main
import ("fmt")

func main() {
	var side_a, side_b int
	fmt.Printf("Enter side a and side b, seperated by whitespaces: ")
	fmt.Scan(&side_a, &side_b)
	fmt.Printf("The result is S = %v and P = %v", (side_a*side_b), (2*(side_a+side_b)))
}