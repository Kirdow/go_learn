package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	Section("Non-Generics", runNonGenerics)
	Section("Generics", runGenerics)
	Section("Any Generics", runAnyGenerics)
	Section("Json Generics", runJsonGenerics)
	Section("Struct Generics", runStructGenercis)
}

func runNonGenerics() {
	var intSlice = []int{1, 2, 3}
	fmt.Printf("Sum Int: %v\n", sumIntSlice(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3}
	fmt.Printf("Sum F32: %v\n", sumFloat32Slice(float32Slice))
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloat32Slice(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}

func runGenerics() {
	var intSlice = []int{1, 2, 3}
	fmt.Printf("Sum Int: %v\n", sumSlice(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3}
	fmt.Printf("Sum F32: %v\n", sumSlice(float32Slice))

	var float64Slice = []float64{1.1, 2.2, 3.3}
	fmt.Printf("Sum F64: %v\n", sumSlice(float64Slice))

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func runAnyGenerics() {
	var intSlice1 = []int{}
	var intSlice2 = []int{1, 2, 3}

	var f32Slice1 = []float32{1, 2, 3}
	var f32Slice2 = []float32{}

	var strSlice1 = []string{"1", "2", "3"}
	var strSlice2 = []string{}

	fmt.Printf("Is Int Empty: %v (expect true)\n", isEmpty(intSlice1))
	fmt.Printf("Is Int Empty: %v (expect false)\n", isEmpty(intSlice2))

	fmt.Printf("Is F32 Empty: %v (expect false)\n", isEmpty(f32Slice1))
	fmt.Printf("Is F32 Empty: %v (expect true)\n", isEmpty(f32Slice2))

	fmt.Printf("Is Str Empty: %v (expect false)\n", isEmpty(strSlice1))
	fmt.Printf("Is Str Empty: %v (expect true)\n", isEmpty(strSlice2))
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func runJsonGenerics() {
	var contacts []contactInfo = loadJSON[contactInfo]("contactInfo")
	fmt.Printf("Contacts : %v\n", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("purchaseInfo")
	fmt.Printf("Purchases: %v\n", purchases)
}

type contactInfo struct {
	Name string
	Email string
}

type purchaseInfo struct {
	Name string
	Price float32
	Amount int
}

func loadJSON[T contactInfo | purchaseInfo](fileName string) []T {
	var filePath = "./json/" + fileName + ".json"

	data, _ := os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

func runStructGenercis() {
	var petrolCar = car[petrolEngine] {
		carMake: "Honda",
		carModel: "Civic",
		engine: petrolEngine{
			litres: 46.94,
			kpl: 17.0,
		},
	}

	var electricCar = car[electricEngine] {
		carMake: "Volkswagen",
		carModel: "ID.7 Pro",
		engine: electricEngine{
			kWh: 96.0,
			kWhp100km: 16.0,
		},
	}

	fmt.Printf("Petrol    Car: %v\n", petrolCar)
	fmt.Printf("Electric  Car: %v\n", electricCar)
}

type petrolEngine struct {
	litres float32
	kpl float32
}

type electricEngine struct {
	kWhp100km float32
	kWh float32
}

type car [T petrolEngine | electricEngine] struct {
	carMake string
	carModel string
	engine T
}
