package main

import (
	"maps"
	"net/http"
	"time"
)

func timeHttpGet(url string) float64 {
	var start = time.Now()

	printfln("Getting URL %s ....", url)
	res, err := http.Get(url)
	panicIf(err)
	defer res.Body.Close()
	panicIfNot(res.StatusCode == http.StatusOK, "expected OK, instead got "+toString(res.StatusCode))

	var elapsed = time.Since(start).Seconds()
	return elapsed
}

func timeHttpGetAndPrint(url string) {
	var time = timeHttpGet(url)
	printfln("Finished getting URL %s, duration of call was:  %f", url, time)
}

func testGoRoutines() {
	go workSimple("A", 1*SECONDS)
	go workSimple("B", 2*SECONDS)
	println("Waiting to finish...")
	time.Sleep(3 * SECONDS)
	mustEqual(3, 3)
}

func testGoRoutineWorkWithChannel() {
	var inChan = make(chan string)
	go work("A", 1*SECONDS, inChan)
	//go workSimple("B", 2*SECONDS)
	//println("Waiting to finish...")
	//time.Sleep(3 * SECONDS)
	var msg = <-inChan
	println("Message = " + msg)
}

func testGoRoutines2NoJoining() {

	var urls = []string{
		"https://www.debian.org",
		"https://www.opensuse.org",
		"https://manjaro.org",
	}

	//	var results = make(map[string]float64)
	for _, url := range urls {
		go timeHttpGetAndPrint(url)
	}

	time.Sleep(3 * SECONDS)
	println("did we get em???")

	// for url, sec := range results {
	// 	printfln("Time for GET of %s is %f seconds", url, sec)
	// }
}

func testReceivingLotsOfMessages() {
	var inChan = make(chan string)
	const doStuffWait = 1 * SECONDS
	go doStuffSendResults(doStuffWait, inChan)
	//go workSimple("B", 2*SECONDS)
	//println("Waiting to finish...")
	//time.Sleep(3 * SECONDS)
	for true {
		println("Waiting for message...")
		msg := <-inChan
		printfln("Received message: %q", msg)
		if msg == doStuffSendResultsLastMessage {
			break
		}
		// const myWait = 4 * SECONDS
		// printfln("Waiting for %+v...", myWait)
		// time.Sleep(myWait)
	}
}

// wait for the first message from either goroutine.
func testSelect() {
	var chanA = make(chan int)
	var chanB = make(chan int)

	go workResult(4, 4*SECOND, chanB)
	go workResult(2, 2*SECOND, chanA)

	var result int
	select {
	case res := <-chanA:
		result = res
	case res := <-chanB:
		result = res
	}

	mustEqual(2, result)
}

func testJoiningWithSelect() {
	var chanA = make(chan int)
	var chanB = make(chan int)
	var chanC = make(chan int)

	const numWorkers = 3
	go workResult(1, 1*SECOND, chanA)
	go workResult(2, 2*SECOND, chanB)
	go workResult(3, 3*SECOND, chanC)

	var results = make(map[chan int]int)
	for len(results) != numWorkers {
		select {
		case res := <-chanA:
			results[chanA] = res
		case res := <-chanB:
			results[chanB] = res
		case res := <-chanC:
			results[chanC] = res
		}
	}

	mustEqual(3, len(results))
	var expected = map[chan int]int{chanA: 1, chanB: 2, chanC: 3}
	mustEqual(true, maps.Equal(expected, results))
}

func testJoiningWithSelectAndTimeout() {
	var chanA = make(chan int)
	var chanB = make(chan int)
	var chanC = make(chan int)

	var results = make(map[chan int]int)
	var done = false

	const numWorkers = 3
	go workResult(1, 1*SECOND, chanA)
	go workResult(2, 2*SECOND, chanB)
	go workResult(3, 3*SECOND, chanC)

	timer := time.NewTimer(2500 * MILLIS)

	for !done && len(results) != numWorkers {
		println("Top of loop")
		select {
		case res := <-chanA:
			results[chanA] = res
		case res := <-chanB:
			results[chanB] = res
		case res := <-chanC:
			results[chanC] = res
		//case <-time.After(2500 * MILLIS): // do not use this; it is not as accurate as time.NewTimer.
		case <-timer.C:
			println("TIMEOUT TRIGGERED")
			done = true
			//break // note: this has no effect inside select
		}
	}
	println("Out of loop.")

	// chanC should be left out!
	mustEqual(2, len(results))
	var expected = map[chan int]int{chanA: 1, chanB: 2}
	mustEqual(true, maps.Equal(expected, results))
}

func testConcurrency() {
	//fmt.Println("Yo")
	//fmt.Printf("Hello, %s!/n", "world")
	//testGoRoutines()
	// testGoRoutines2NoJoining()
	//testGoRoutineWorkWithChannel()
	//testReceivingLotsOfMessages()
	//testSelect()
	//testJoiningWithSelect()
	testJoiningWithSelectAndTimeout()
}
