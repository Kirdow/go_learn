package main

import "fmt"

var firstSection bool = true
func Section(title string, fn func()) {
	if !firstSection {
		var i = 2
		for i > 0 {
			i -= 1
			fmt.Println()
		}
	}

	fmt.Printf("] >> %s\n", title)
	firstSection = false

	fn()
}
