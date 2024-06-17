package main

import "fmt"

// Engine interface defines what an engine should do
type Engine interface {
	Start()
	Stop()
}

// GasolineEngine implements the Engine interface
type GasolineEngine struct{}

func (ge *GasolineEngine) Start() {
	fmt.Println("Gasoline engine starting...")
}

func (ge *GasolineEngine) Stop() {
	fmt.Println("Gasoline engine stopping...")
}

// ElectricEngine implements the Engine interface
type ElectricEngine struct{}

func (ee *ElectricEngine) Start() {
	fmt.Println("Electric engine starting...")
}

func (ee *ElectricEngine) Stop() {
	fmt.Println("Electric engine stopping...")
}

// Car struct has an Engine dependency
type Car struct {
	engine Engine
}

func (c *Car) Drive() {
	c.engine.Start()
	fmt.Println("Car is driving...")
	c.engine.Stop()
}

func main() {
	// Create a car with a gasoline engine
	gasolineCar := &Car{engine: &GasolineEngine{}}
	gasolineCar.Drive()

	// Create a car with an electric engine
	electricCar := &Car{engine: &ElectricEngine{}}
	electricCar.Drive()
}

