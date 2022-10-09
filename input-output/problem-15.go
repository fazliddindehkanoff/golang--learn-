// this solution for calculating diametr and radius of circle by it's area
package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func main(){
	var area float64
	fmt.Printf("Enter the area of circle: ")
	fmt.Scan(&area)
	fmt.Printf("Diametr of Circle: %v\nRadius of Circle: %v", (2*(math.Sqrt((area/Pi)))), (math.Sqrt((area/Pi))))
}
