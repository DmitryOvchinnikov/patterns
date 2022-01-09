package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Printf("\nMaking an Americano\n----------------\n")

	americano := &CoffeeMachine{}
	beans := (size / 8.0) * 5.0
	americano.startCoffee(beans, 2)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()

	fmt.Println("Americano is ready!")
}

func makeLatte(size float32, foam bool) {
	fmt.Printf("\nMaking an Latte\n----------------\n")

	latte := &CoffeeMachine{}
	beans := (size / 8.0) * 5.0
	latte.startCoffee(beans, 4)
	latte.grindBeans()
	latte.useHotWater(size)

	milk := (size / 8.0) * 2.0
	latte.useMilk(milk)
	latte.doFoam(true)
	latte.endCoffee()

	fmt.Println("Latte is ready!")
}
