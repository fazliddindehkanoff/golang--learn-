package main
import ("fmt")

const Pi = 3.14

func main(){
	var length_of_circle, radius float64
	fmt.Printf("Enter the length of Circle: ")
	fmt.Scan(&length_of_circle)
	radius = length_of_circle/(2*Pi)
	fmt.Printf("Radius of Circle: %v\nArea of Circle: %v", radius, (Pi*(radius*radius)))
}
