package main

import "fmt"

func main() {
	var myAge float64
	fmt.Println("What is your age, dear User?")
	fmt.Scan(&myAge)
	fmt.Printf("All right! So, you are %.2f years old.", myAge)
}
