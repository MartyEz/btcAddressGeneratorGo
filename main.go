package main

import (
	"btcAddressGeneratorGo/generator"
	"btcAddressGeneratorGo/utils"
	"fmt"
	"sync"
	"time"
)

// Number of goroutine
var ROUTINE_NUMBER int = 500


func main() {

	// Create chan String for each goRoutine
	chanArr := make([]chan []byte, ROUTINE_NUMBER)

	var waitGroup sync.WaitGroup
	start := time.Now()

	// Launch routines with a unique chan binded
	for i := 0; i < ROUTINE_NUMBER; i++ {
		chanArr[i] = make(chan []byte, 100000)
		go generator.GenerateAdr(&waitGroup, chanArr[i])
	}

	workSended := false
	y := 0

	for i := 0; i < 5000; i++ {
		waitGroup.Add(1)
		s := []byte(utils.RndString()) // generate random string
		workSended = false

		// Check size's of routine's work queue and send work to not full queue
		for !workSended {
			if len(chanArr[y]) < cap(chanArr[y]) {
				chanArr[y] <- s
				workSended = true
			}
			y = (y + 1)% ROUTINE_NUMBER
		}

	}

	// Close all string chan
	for i := 0; i < ROUTINE_NUMBER; i++ {
		close(chanArr[i])
	}

	// wait that all routine
	waitGroup.Wait()

	duration := time.Since(start)
	fmt.Println("File Tested in :", duration.Seconds())
}
