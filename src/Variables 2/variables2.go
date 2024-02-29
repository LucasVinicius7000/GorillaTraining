package main

import "fmt"

func main() {
	var (
		peopleAge [5]int
		average   float64
	)
	for i := 0; i < 5; i++ {
		fmt.Printf("\nHow old is the %dº people?\n", i+1)
		fmt.Scan(&peopleAge[i])
		average += float64(peopleAge[i])
	}
	average /= float64(len(peopleAge))
	fmt.Printf("\n The average of people age is %.2f\n", average)

}
