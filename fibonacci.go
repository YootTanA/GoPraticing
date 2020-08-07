package main

import "fmt"

func fibonacci() func() int {

	firstDigit := 0
	secondDigit := 1
	i := -1
	sum := 0

	return func() int {

		i++
		if i > 1 {

			sum = firstDigit + secondDigit
			firstDigit = secondDigit
			secondDigit = sum

			return sum

		}

		if i == 0 {

			return 0

		} else {

			return 1

		}

	}
}

func main() {

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
