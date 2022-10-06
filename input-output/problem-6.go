//This solution for calculating parallelepiped's area and volume
package main
import ("fmt")

func main() {
	var side_a, side_b, side_c int
	fmt.Printf("Enter parallelepiped's sides a, b, c (seperated with whitespace): ")
	fmt.Scan(&side_a, &side_b, &side_c)
	fmt.Printf("Parallelepiped's Volume is equal to: %v Area is equal to: %v", (side_a*side_b*side_c), (2*(side_a*side_b+side_a*side_c+side_b*side_c)))
}