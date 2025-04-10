package main

import "fmt"

// to create a strucr you can use the type keyword
type gasEngine struct {
	milesPerGallon uint32
	gallon         uint32
	ownerInfo      owner
}
type electricEngine struct {
	milesPerKilloWattHour uint32
	killoWattHour         uint32
	ownerInfo             owner
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint32
}

// you can also bind methods to structs
// aloha is a method bound to gasEngine

func (e gasEngine) milesLeft() uint32 {
	return e.milesPerGallon * e.gallon
}
func (e electricEngine) milesLeft() uint32 {
	return e.milesPerKilloWattHour * e.killoWattHour
}

func canMakeIt(e engine) bool {
	if e.milesLeft() < 1000 {
		return true
	} else {
		return false
	}
}

func structsInterfaces() {
	var myEngine gasEngine // this is a "zero" valued struct

	var anotherEngine gasEngine = gasEngine{milesPerGallon: 32, gallon: 10, ownerInfo: owner{name: "reesav"}} //or you can do this syntax

	var anotherEngine2 gasEngine = gasEngine{32, 10, owner{"reesav"}} // or this aswell

	// you can directly set the fields as well
	myEngine.milesPerGallon = 10
	myEngine.gallon = 10
	anotherEngine2.milesPerGallon = 10

	fmt.Printf("\nthis is the struct:%v", anotherEngine)

	milesLeft := myEngine.milesLeft()

	fmt.Printf("\nthis is milesLeft: %v", milesLeft)

	var elecEng electricEngine = electricEngine{milesPerKilloWattHour: 20, killoWattHour: 22, ownerInfo: owner{name: "reesav"}}

	if canMakeIt(elecEng) {
		fmt.Printf("this is cool")
	} else {
		fmt.Printf("not cool")
	}

}
