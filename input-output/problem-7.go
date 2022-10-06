// This solution for calculating circle's area and length by circle's Radius
package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func main() {
	var radius float64
	fmt.Printf("Enter Circle's Radius: ")
	fmt.Scan(&radius)
	fmt.Printf("Circle's length(L): %v\nCircle's Area(S): %v", (2*Pi*radius),
	(Pi*math.Pow(radius, 2)))
}