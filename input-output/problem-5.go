// this solution for solving cube's volume and surface are
package main

import (
	"fmt"
	"math"
)

func main() {
	var a_side float64
	fmt.Printf("Enter a side of cube: ")
	fmt.Scan(&a_side)
	fmt.Printf("Volume is: %v, Surface are is: %v", (math.Pow(a_side, 3)), (6*math.Pow(a_side, 2)))
}