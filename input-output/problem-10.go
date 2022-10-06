package main
import ("fmt")

func main() {
	var first_num, second_num float32
	fmt.Printf("Enter 2 numbers (seperated by whitespaces): ")
	fmt.Scan(&first_num, &second_num)
	fmt.Printf("Results: addition of nums: %v\nMultiplication of them: %v\nSquare of %v = %v\nSquare of %v = %v", (first_num + second_num),(first_num * second_num), first_num, (first_num*first_num), second_num, (second_num*second_num))
}