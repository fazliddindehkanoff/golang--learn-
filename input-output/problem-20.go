package main
import (
	"fmt"
	"math"
)


func main() {
	var x1, x2, y1, y2 float64
	fmt.Printf("Enter 4 numbers (seperated by whitespaces): ")
	fmt.Scan(&x1, &x2, &y1, &y2)
	fmt.Printf("The result is: %v", (math.Sqrt(math.Pow((x2-x1),2)+math.Pow((y2-y1), 2))))
}