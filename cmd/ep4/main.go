package main

import (
	"fmt"
	"time"
)

func main() {
	Section("Arrays && Slices", runArraysAndSlices)
	Section("Maps", runMaps)
	Section("Time Slices", runTimeSlices)
}

func runArraysAndSlices() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Printf("IntArr[0] = %d\n", intArr[0])
	fmt.Printf("IntArr[1:3] = %v\n", intArr[0:3])

	fmt.Printf("&IntArr[0] = %v\n", &intArr[0])
	fmt.Printf("&IntArr[1] = %v\n", &intArr[1])
	fmt.Printf("&IntArr[2] = %v\n", &intArr[2])

	var intArr2 [3]int32 = [3]int32{1,2,3}
	fmt.Printf("IntArr2 [Inline] = %v\n", intArr2)

	intArr3 := [3]int32{4,5,6}
	fmt.Printf("IntArr3 [Inline+Infer] = %v\n", intArr3)

	intArr4 := [...]int32{7,8,9}
	fmt.Printf("IntArr4 [Inline+Infer(Type+Len)] = %v\n", intArr4)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("IntSlice: %v\n", intSlice)
	fmt.Printf("Len = %d, Cap = %d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("IntSlice [Append(7)]: %v\n", intSlice)
	fmt.Printf("Len = %d, Cap = %d\n", len(intSlice), cap(intSlice))

	fmt.Println("CAP DEMO START")
	var testSlice []int32 = []int32{1, 2, 3}
	var i int32 = 4

	var lastCap int = 0
	// Technically too early for this, but hey, sometimes that's what we do
	for i <= 200 {
		var newCap int = cap(testSlice)
		if newCap != lastCap {
			fmt.Printf("Cap = %d -> %d, Len = %d\n", lastCap, newCap, len(testSlice))
			lastCap = newCap
		}

		testSlice = append(testSlice, i)
		i += 1
	}

	fmt.Printf("Last Cap = %d, Len = %d\n", cap(testSlice), len(testSlice))
	fmt.Println("CAP DEMO END")

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("IntSlice [Spread]: %v\n", intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("IntSlice3 = %v, Len = %d, Cap = %d\n", intSlice3, len(intSlice3), cap(intSlice3))
}

func printAge(ageMap map[string]uint8, name string) bool {
	var age, ok = ageMap[name]

	if ok {
		fmt.Printf("%s is %dyo\n", name, age)
	} else {
		fmt.Printf("We don't know of %s\n", name)
	}

	return ok
}

func runMaps() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Printf("MyMap: %v\n", myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Printf("MyMap2: %v\n", myMap2)
	fmt.Println("Direct Approach")
	fmt.Printf("Adam Age: %d\n", myMap2["Adam"])
	fmt.Printf("Jason Age: %d\n", myMap2["Jason"])
	fmt.Println("Check Approach")
	printAge(myMap2, "Adam")
	printAge(myMap2, "Jason")

	fmt.Println("Add Jason, Remove Adam")
	myMap2["Jason"] = 38
	delete(myMap2, "Adam")

	printAge(myMap2, "Adam")
	printAge(myMap2, "Jason")

	fmt.Printf("MyMap2: %v\n", myMap2)

	fmt.Println("Add Names")

	myMap2["Emil"] = 19
	myMap2["Joana"] = 33

	fmt.Println("Iterate Names:")
	for name := range myMap2 {
		fmt.Printf("Name: %s\n", name)
	}

	fmt.Println("Iterate Names and Ages:")
	for name, age := range myMap2 {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}

func runTimeSlices() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var timeStart = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(timeStart)
}
