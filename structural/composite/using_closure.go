package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming")
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

type SharkA struct {
	Animal
	MySwim func()
}
