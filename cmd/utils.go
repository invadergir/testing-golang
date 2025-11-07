package main

import (
	"fmt"
	"strconv"
	"time"
)

func printlnf(format string, a ...any) {
	fmt.Printf(format, a...)
	fmt.Println("")
}

func printfln(format string, a ...any) {
	printlnf(format, a...)
}

func panicIf(err any) {
	if err != nil {
		fmt.Println("==============================================")
		panic(err)
	}
}

func panicIfNot(condition bool, error string) {
	if !condition {
		fmt.Println("==============================================")
		panic(error)
	}
}

const SECOND = time.Second
const SECONDS = time.Second
const MILLI = time.Millisecond
const MILLIS = time.Millisecond

func workSimple(id string, timeout time.Duration) {
	time.Sleep(timeout)
	fmt.Printf("Work('%s', %v) done.\n", id, timeout)
}

// work func that signals when it's done.
func work(id string, timeout time.Duration, outChan chan<- string) {
	time.Sleep(timeout)
	var result string = fmt.Sprintf("Work('%s', %v) done.\n", id, timeout)

	outChan <- result
}

// return the given result after the specified time.
func workResult(result int, timeout time.Duration, outChan chan<- int) {
	printfln("workResult(%d, %v) working...", result, timeout)
	time.Sleep(timeout)
	printfln("workResult(%d, %v) DONE!", result, timeout)
	outChan <- result
}

const doStuffSendResultsLastMessage = "DONE"

func doStuffSendResults(timeout time.Duration, outChan chan<- string) {
	const numIter = 5
	for ix := range numIter {
		time.Sleep(timeout)
		var result = fmt.Sprintf("Did stuff.  Result is %d", ix)
		outChan <- result
	}
	outChan <- doStuffSendResultsLastMessage
}

func toString(i int) string {
	return strconv.Itoa(i)
}

func mustEqual[T comparable](expected T, actual T) {
	if expected != actual {
		panic(fmt.Errorf("We expected to get '%+v' but instead got '%+v'\n", expected, actual))
	}
}

func mustNotEqual[T comparable](expected T, actual T) {
	if expected == actual {
		panic(fmt.Errorf("We expected to NOT get '%+v' but instead got '%+v'\n", expected, actual))
	}
}

// func workFor2Sec(id string) {
// 	time.Sleep(2 * SECONDS)
// 	fmt.Printf("WorkFor2Sec('%s') done.", id)
// }
