package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	Section("Channels", runChannels)
	Section("Fib", runFib)
	Section("Buffer", runBuffer)
	Section("Sale Price", runSalePrice)
}

func runChannels() {
	var c = make(chan int)
	go process(c)
	var i = <- c
	fmt.Printf("Channel: %d\n", i)
}

func process(c chan int) {
	c <- 123
}

func runFib() {
	var c = make(chan int)
	go generate(c, 10)
	for i := range c {
		fmt.Printf("Next: %d\n", i)
	}
}

func generate(pipe chan int, iter int) {
	defer close(pipe)

	var a int = 0
	var b int = 1
	var i int = 0

	for i < iter {
		i++

		var c = a + b
		pipe <- c

		a = b
		b = c
	}
}

func runBuffer() {
	var c = make(chan int, 5)
	go generate2(c, 10)
	for i := range c {
		fmt.Printf("Next: %d\n", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func generate2(pipe chan int, iter int) {
	defer close(pipe)

	var a int = 0
	var b int = 1
	var i int = 0

	for i < iter {
		i++

		var c = a + b
		pipe <- c

		a = b
		b = c
	}
	
	fmt.Printf("Exiting generate2\n")
}

var MAX_CHICKEN_PRICE float32 = 50
var MAX_TOFU_PRICE float32 = 30

func runSalePrice() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"ica.se", "willys.se", "coop.se"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 200
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 200
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <- chickenChannel:
		fmt.Printf("Text Sent: Found a deal on chicken at %s\n", website)
	case website := <- tofuChannel:
		fmt.Printf("Email Sent: Found a deal on tofu at %s\n", website)
	}
}
