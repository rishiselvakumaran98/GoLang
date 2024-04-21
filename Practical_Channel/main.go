package main

import (
	"fmt"
	"time"
	"math/rand"
)
 
var MAX_VEGTABLE_PRICE float32 = 7
var MAX_TOFU_PRICE float32 = 7
// lets construct a Vegatble price querying tool
func main(){
	var vegChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com","costco.com", "samsclub.com"}
	for _, website := range websites{
		go vegCheckPrices(website, vegChannel)
		go tofuCheckPrices(website, tofuChannel)
	}
	sendMessage(vegChannel, tofuChannel)
}

func vegCheckPrices(website string, vegChannel chan string){
	for {
		time.Sleep(time.Second*1)
		var vegPrice = rand.Float32()*20
		if vegPrice <= MAX_VEGTABLE_PRICE{
			vegChannel <- website
			break
		}
	}
}

func tofuCheckPrices(website string, tofuChannel chan string){
	for {
		time.Sleep(time.Second*1)
		var tofuPrice = rand.Float32()*20
		if tofuPrice <= MAX_TOFU_PRICE{
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(vegChannel chan string, tofuChannel chan string){
	// lets say we are also having a tofuPrice check then we listen for price from both Channel
	select {
		case website := <- vegChannel:
			fmt.Printf("\nFound a deal for a vegetable at %v\n", website)
		case website := <- tofuChannel:
			fmt.Printf("\nFound a deal for a Tofu at %v\n", website)
	}
	
}