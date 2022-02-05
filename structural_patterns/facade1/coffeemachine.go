package main

import "fmt"

// CoffeeMachine represents an API to a hypothetical coffee maker.
type CoffeeMachine struct {
	beanAmount   float32
	grinderLevel int
	waterTemp    int
	waterAmount  float32
	milkAmount   float32
	addFoam      bool
}

func (c *CoffeeMachine) startCoffee(beans float32, grind int) {
	c.beanAmount = beans
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beans:", beans, "and grind level:", grind)
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans:", c.beanAmount, "beans at", c.grinderLevel)

	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount, "oz")
	c.milkAmount = amount

	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temp:", temp)
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water:", amount)
	c.waterAmount = amount

	return amount
}

func (c *CoffeeMachine) doFoam(foam bool) {
	fmt.Println("Foam setting:", foam)
	c.addFoam = foam
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order.")
}
