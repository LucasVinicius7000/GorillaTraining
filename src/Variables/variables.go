package main

import "fmt"

func main() {
	fmt.Println("What is your age, dear User?")
	var (
		age        int
		yearBirth  int
		monthBirth int
		dayBirth   int
	)
	fmt.Scan(&age)
	fmt.Println("Sure, what year were you born?")
	fmt.Scan(&yearBirth)
	fmt.Println("Nice, which month?")
	fmt.Scan(&monthBirth)
	fmt.Println("and.. the day?")
	fmt.Scan(&dayBirth)
	fmt.Print("\n\nCool!\n")
	currentYear := yearBirth + age
	fmt.Printf("So.. based on this info i think the current year is %d\n\n\n", currentYear)
}
