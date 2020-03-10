package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if isEven(number) {
			fmt.Printf("%v is Even\n", number)
		} else {
			fmt.Printf("%v is Odd\n", number)
		}
	}
}

func isEven(number int) bool {
	if number % 2 == 0 {
		return true
	}

	return false
}
