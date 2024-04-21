package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
var m = sync.RWMutex{} // We use Mutual Exclusion to lock/UnLock the threads
var wg = sync.WaitGroup{} // Acts like a counter where each task adds a 1 token
// need to wait gor the counter to go to 0 to mark as done

var dbData = []string{"id1", "id2","id3","id4","id5"}
var results = []string{}
func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("\nTotal executed time: %v", time.Since(t0))
	fmt.Println("\nThe value of results is: %v", results)
}

func dbCall(i int){
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond) // time.Duration(delay)*time.Millisecond prints as 2ms
	fmt.Println("Printing from DB: ", dbData[i])
	save(dbData[i])
	log()
	//m.Lock() // if a lock is placed the methods which reach here check to see if the lock is unlocked before proceeding, which helps to essentially "queue" to use the code implementation
	//results = append(results, dbData[i]) // this can lead to concurrency issues since we have multiple threads writing into the results function at the same time
	//m.Unlock()
	wg.Done()
}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Println("Printing results: ", results)
	m.RUnlock()
}