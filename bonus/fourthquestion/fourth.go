package fourthquestion

import "fmt"

// EgyptVersionCalculate калькулирует по египетски))
func EgyptVersionCalculate(digit1, digit2 int) {
	if digit2 == 0 || digit1 == 0 {
		return
	}

	if digit2%digit1 == 0 {
		fmt.Print("1/", digit2/digit1)
		return
	}

	if digit1%digit2 == 0 {
		fmt.Print(digit1 / digit2)
		return
	}

	if digit1 > digit2 {
		fmt.Print(digit1/digit2, " + ")
		EgyptVersionCalculate(digit1%digit2, digit2)
		return
	}

	n := digit2/digit1 + 1
	fmt.Print("1/", n, " + ")

	EgyptVersionCalculate(digit1*n-digit2, digit2*n)
}
