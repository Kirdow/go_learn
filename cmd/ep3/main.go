package main

import (
	"fmt"
	"errors"
)

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)

	var num int = 17
	var den int = 5
	var result int = intDivision(num, den)
	fmt.Printf("[Int Div] %d / %d = %d\n", num, den, result)

	var frac1, rem1 = intDivision2(num, den)
	fmt.Printf("[Int Div+Rem] A=%d, B=%d, (A / B, A %% B) = (%d, %d)\n", num, den, frac1, rem1)

	runDiv3(num, den)
	runDiv3(num, 0)
	runDiv3(den * num, den)
}

func runDiv3(num int, den int) {
	var frac, rem, err = intDivOrError(num, den)
	switch {
	case err != nil:
		fmt.Printf("[Error Compat] A=%d, B=%d, (A / B), Error: %s\n", num, den, err.Error())
		return
	case rem == 0:
		fmt.Printf("[Error Compat] A=%d, B=%d, (A / B) = %d\n", num, den, frac)
	default:
		fmt.Printf("[Error Compat] A=%d, B=%d, (A / B, A %% B) = (%d, %d)\n", num, den, frac, rem)
	}

	switch rem {
		case 0:
			fmt.Println("The division was exact")
		case 1,2:
			fmt.Println("The division was close")
		default:
			fmt.Println("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, den int) int {
	var result int = num / den
	return result
}

func intDivision2(num int, den int) (int, int) {
	var frac int = num / den
	var rem int = num % den

	return frac, rem
}

func intDivOrError(num int, den int) (int, int, error) {
	var err error
	if den == 0 {
		err = errors.New("Cannot Divide By Zero")
		return 0, 0, err
	}
	var frac int = num / den
	var rem int = num % den

	return frac, rem, err
}
