package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}


type owner struct{
	name string
}

func (e gasEngine) calculateMilesLeft() uint8 { // this function belongs to the gasEngine class
	return e.gallons * e.mpg
}
// What if we wanted to make our calculateMilesLeft more generic to support
// both gasEngine and electricEngine? We use interfaces for that
func (e electricEngine) calculateMilesLeft() uint8 { // this function belongs to the electricEngine class
	return e.mpkwh * e.kwh
}

type Engine interface{
	calculateMilesLeft() uint8
}

func canMakeIt(e Engine, miles uint16) bool {
	return  uint16(miles) <= uint16(e.calculateMilesLeft())
}

func main(){
	// structs are a way of defining our own type in Go
	// if not initialized they will be set to default values 
	// The struct below is reusable
	var elantraEngine gasEngine = gasEngine{mpg:40, gallons: 12, owner: owner{"Alex"}} // can also be set as gasEngine{40, 12}
	var distanceToTravel uint16 = 300
	fmt.Printf("Miles left on elantraEngine: %v, Distance to travel: %v\n", elantraEngine.calculateMilesLeft(), distanceToTravel)
	
	var teslaEngine electricEngine = electricEngine{mpkwh: 20, kwh: 10}
	fmt.Printf("Miles left on teslaEngine: %v, Distance to travel: %v\n", teslaEngine.calculateMilesLeft(), distanceToTravel)
	
	if canMakeIt(elantraEngine, distanceToTravel){
		fmt.Println("Yes my car can make!!")
	} else {
		fmt.Println("Damn I am doomed my car cannot make!!")
	}
	
	// The struct below is "not" reusable -> meaning it has to be redefined as a struct
	var elantraEngine2 = struct{
		mpg uint8
		gallons uint8
	}{25, 15}
	fmt.Printf("Reusable struct %v vs Not-reusable Struct %v\n", elantraEngine, elantraEngine2)
}