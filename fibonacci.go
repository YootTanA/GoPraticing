package main

import "fmt"

func fibonacci() func() int {

	firstDigit := 0
	secondDigit := 1
	sum := -1

	return func() int {

		sum++
		if sum > 1 {

			sum = firstDigit + secondDigit
			firstDigit = secondDigit
			secondDigit = sum

			return sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
