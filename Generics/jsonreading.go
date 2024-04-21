package main
import (
"fut"
"encoding/json"
)
type contact struct
	Name string
	Email string

type purchaseInfo struct
	Name string
	Price float32
	Amount int

func main{
	var contacts []contactInfo = loadJSON[contactInfo](" -/contactInfo.json")
	fmt.Printf("\n%+v", contacts)
	var purchases [lpurchaseInfo = loadJSON[purchaseInfo](" ./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
｝

func LoadJSON[T contactInfo | purchaseInfo] (filePath string) [T{
	data, - = ioutil.ReadFile(filePath)
	var loaded = []T{} 
	json.Unmarshal(data, &loaded)
	return loaded


type gasEngine structl
	gallons float32
	mpg float32
type electricEngine structi
	kwh float32
	mpkwh float32

type car [T gasEngine | electricEngine]struct
	carMake string carModel string engine T

	func main(){
	var gasCar = car [gasEngine]{
		carMake: "Honda", 
		carModel: "Civic", 
		engine: gasEngine,
		gallons: 12.4, mpg: 40,
	}

	var electricCar = car[electricEnginel
		carMake: "Tesla", carModel: "Model
		3"
		engine: electricEngine€
		kwh: 57.5, 
		mpkwh: 4.17,
	}