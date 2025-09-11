package main

import "fmt"

func main() {
	Section("Pointers", runPointers)
	Section("Pointers w/ Functions", runPointersWithFunctions)
}

func runPointers () {
	// *typeName is pointer-type
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("*P = %v\n", *p)
	fmt.Printf("I = %v\n", i)

	// *varName is de-reference
	*p = 10
	i = 7

	fmt.Printf("*P = %v\n", *p)
	fmt.Printf("I = %v\n", i)
	
	// Copy the address of i
	p = &i

	fmt.Printf("*P = %v\n", *p)
	fmt.Printf("I = %v\n", i)
}

func runPointersWithFunctions() {
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("&Thing1 = %p\n", &thing1)

	var result [5]float64 = square(thing1)
	fmt.Printf("Result = %v\n", result)
	fmt.Printf("Thing1 = %v\n", thing1)

	var result2 [5]float64 = squareRef(&thing1)
	fmt.Printf("Result = %v\n", result2)
	fmt.Printf("Thing1 = %v\n", thing1)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("&Thing2 = %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}

func squareRef(thing3 *[5]float64) [5]float64 {
	fmt.Printf("&Thing3 = %p\n", &thing3)
	for i := range thing3 {
		thing3[i] = thing3[i]*thing3[i]
	}

	return *thing3
}

 
