package main

import "fmt"

import "unicode/utf8"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Printf("IntNum + 1 = %d\n", intNum)

	var floatNum float64 = 12345678.9
	fmt.Printf("Float64 = %.2f\n", floatNum)

	var floatA_32 float32 = 0.1
	var floatB_32 float32 = 0.2
	var floatC_32 float32 = floatA_32 + floatB_32
	fmt.Printf("%.8f + %.8f = %.10f\n", floatA_32, floatB_32, floatC_32)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Printf("float[%.2f] + cast<float>(int[%d]) = float[%.2f]\n", floatNum32, intNum32, result)

	var intNum1 int = 7
	var intNum2 int = 5
	fmt.Printf("7 / 5 = %d\n", intNum1 / intNum2);
	fmt.Printf("7 %% 5 = %d\n", intNum1 % intNum2);

	var myString string = "Hello, World!"
	fmt.Println(myString)

	var myTwoLineString string = `Hello,
	World!`

	fmt.Println(myTwoLineString)

	fmt.Printf("Byte Len A: %d\n", len("A"))
	fmt.Printf("Byte Len [Gamma(γ)]: %d\n", len("γ"))
	fmt.Printf("UTF8 Len A: %d\n", utf8.RuneCountInString("A"))
	fmt.Printf("UTF8 Len [Gamma(γ)]: %d\n", utf8.RuneCountInString("γ"))

	var myRune rune = 'a'
	fmt.Printf("My Rune: %c\n", myRune)
	fmt.Println(myRune);

	var myBoolean bool = false
	fmt.Printf("My Bool: %t\n", myBoolean)

	var intNum3 int
	fmt.Printf("Default Int: %d\n", intNum3)

	myVar := "text"
	fmt.Printf("ShortHand: %s\n", myVar)

	var1, var2 := 1, 2
	fmt.Printf("My short-hand vars: ")
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Printf("My Const: %s\n", myConst)

	const piStr string = "3.141592653589793238462643383"
	const pi32 float32 = 3.141592653589793238462643383
	const pi64 float64 = 3.141592653589793238462643383
	fmt.Println("             Wrong --v");
	fmt.Printf("Const PI 32: %.27f\n", pi32)
	fmt.Printf("Correct  PI: %s\n", piStr)
	fmt.Printf("Const PI 64: %.27f\n", pi64)
	fmt.Println("                      Wrong --^");
}
