package main

import "fmt"

type gasEngine struct {
	mileage  uint16
	capacity uint16
	owner
}

type electricEngine struct {
	dpc    uint16
	charge uint16
	owner
}

type owner struct {
	name string
}

func (e gasEngine) distance() uint16 {
	return e.capacity * e.mileage
}

func (e electricEngine) distance() uint16 {
	return e.dpc * e.charge
}

type engine interface {
	distance() uint16
}

func canmakeit(e engine, km uint16) {
	if km <= e.distance() {
		fmt.Println("go boy")
	} else {
		fmt.Println("dont ")
	}
}

func main() {
	owner1 := owner{"Salih"}
	myEngine := gasEngine{20, 30, owner1}
	myEngine.capacity = 40
	fmt.Println(myEngine.mileage, myEngine.capacity, myEngine.owner)
	prajol := owner{"prajol"}
	prajolengine := electricEngine{4, 65, prajol}
	fmt.Println("this is so grand i10")
	canmakeit(myEngine, 255)
	fmt.Println("\n this is so gay")
	canmakeit(prajolengine, 5000)
}
