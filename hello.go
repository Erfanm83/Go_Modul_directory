package main

import (
	"fmt"
	"math"
)

func main() {
	// i ,err := strconv.Atoi("-42")
	// s := strconv.Itoa(-42)
	// b , err := strconv.ParseBool("true")
	// f , err := strconv.ParsFloat("3.1" , 64)
	// i , err := strconv.ParsInt("3.1" , 10 , 64)
	// u , err := strconv.ParseUnit("3.1" ,10 , 64)

	var a, b, c float64
	var delta, answer1, answer2 float64
	fmt.Scan(&a, &b, &c)

	if a < 0 {
		a = -a
	}

	if a == 0 {
		answer1 = -c / b
		fmt.Printf("%.3f", answer1)
		if b == 0 {
			fmt.Print("IMPOSSIBLE")
		}
	} else {

		delta = (b * b) - (4 * a * c)

		if delta < 0 {
			fmt.Print("IMPOSSIBLE")
		} else if delta > 0 {
			answer1 = -b + math.Sqrt(delta)/(2*a)
			fmt.Println("%.3f", answer1)
			answer2 = (-b - math.Sqrt(delta)) / (2 * a)
			fmt.Printf("%.3f", answer2)
		} else if delta == 0 {
			answer2 = (-b - math.Sqrt(delta)) / (2 * a)
			fmt.Printf("%.3f", answer2)
		}
	}
}
