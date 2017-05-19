package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// argumenter
	if 1 < len(os.Args) {
		fmt.Print(len(os.Args) - 1)
		fmt.Println(" arguments provided")
	}

	//Scanf venter på en input
	fmt.Println("Enter first prime number")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("not a number")
	} else {

		IsPrimeSqrt(i)

	}

	fmt.Println("Enter second prime number")
	var x int
	_, err2 := fmt.Scanf("%d", &x)
	if err2 != nil {
		fmt.Println("not a number")
	} else {

		IsPrimeSqrt(x)

	}

}
func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			fmt.Println(value, " is not a prime number ")
			return false
		}
	}
	fmt.Println(value, "is a prime number")

	return value > 1
}
