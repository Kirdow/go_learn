package main

import "fmt"

func main() {
	Section("Structs && Interfaces", runStructsAndInterfaces)
}

type petrolEngine struct {
	kpl uint16
	litres uint16
}

type electricEngine struct {
	kwhp100km uint16
	kwh uint16
}

func (e petrolEngine) distanceLeft() uint16 {
	return e.litres * e.kpl
}

func (e electricEngine) distanceLeft() uint16 {
	var mpkwh uint32 = 100000 / uint32(e.kwhp100km)
	return uint16(mpkwh * uint32(e.kwh) / 1000)
}

type engine interface {
	distanceLeft() uint16
}

func canMakeIt(e engine, distance uint16) {
	if distance <= e.distanceLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to refuel first!")
	}
}

func drive(e engine, distance uint16) {
	fmt.Printf("You want to drive %dkm\n", distance)
	canMakeIt(e, distance)
}

func runStructsAndInterfaces() {
	// can also provide keys by doing = petrolEngine{kpl: 10, litres: 56}
	var myEngine petrolEngine = petrolEngine{10, 56}
	fmt.Printf("MyEngine[%v %v]\n", myEngine.kpl, myEngine.litres)
	fmt.Printf("Distance left: %dkm\n", myEngine.distanceLeft())

	drive(myEngine, 500)
	drive(myEngine, 600)

	var myElectricEngine electricEngine = electricEngine{17, 96}
	fmt.Printf("MyElectricEngine[%v %v]\n", myElectricEngine.kwhp100km, myElectricEngine.kwh)
	fmt.Printf("Distance left: %dkm\n", myElectricEngine.distanceLeft())

	drive(myElectricEngine, 500)
	drive(myElectricEngine, 600)

}

 
