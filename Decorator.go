package main

import "fmt"

func main() {
	espresso := &Espresso{}

	espressoWithMilk := &Whip{espresso}

	fmt.Println("Price of espresso with whip is ", espressoWithMilk.cost())
}

type Beverage interface {
	cost() float32
}

type Espresso struct {
}

type DarkRoast struct {
}

type Milk struct {
	beverage Beverage
}

type Whip struct {
	beverage Beverage
}

func (c *Espresso) cost() float32 {
	return 1
}

func (c *DarkRoast) cost() float32 {
	return 1
}

func (t *Milk) cost() float32 {
	coffeePrice := t.beverage.cost()
	return coffeePrice + 0.2
}

func (t *Whip) cost() float32 {
	coffeePrice := t.beverage.cost()
	return coffeePrice + 0.25
}
