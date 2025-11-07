package main

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime/debug"
	"strings"
	"testing"
)

// This is very similar to the 'testify' unit testing framework package
// EXCEPT that you can create an instance of 'asserter' and then call it
// without passing in 't' every time.
// Only limited assertions are implemented currently.

type Asserter struct {
	testingT *testing.T
}

func MakeAsserter(t *testing.T) *Asserter {
	return &Asserter{t}
}

// func (t *tester) equals[C comparable](actual C, expected C) bool {

func (a *Asserter) PrintDivider() {
	fmt.Println("=========================================================")
}

// numEntriesToSkip = the number of stack entries to "skip", meaning filter out.
//
//	We want to print only the line where the assertion is called.
//	So if you call this from an assertion, give 1.
//	If your code is one level above an assertion, give 2.
//	If you want to see the whole stack trace, give 0.
func (a *Asserter) GetRelevantStack(numBackUp int) []string {

	// fmt.Println("\\\\\\\\\\\\\\")
	// fmt.Println("BEFORE stack is: ")
	// debug.PrintStack()
	// fmt.Println("///////")

	// each 'entry' is two lines
	const lengthOfEntry = 2
	var printFromHere = 0
	if numBackUp > 0 {
		printFromHere = numBackUp*lengthOfEntry + lengthOfEntry + 1
	}
	var stack = string(debug.Stack())
	var stackLines = strings.Split(stack, "\n")
	if printFromHere == 0 {
		return stackLines
	} else {
		var shortStack = make([]string, 0, lengthOfEntry)
		var numBackedUp = 0
		var thisFuncEntryRE = regexp.MustCompile("^[^ \t].*\\(\\*asserter\\)\\.GetRelevantStack\\(0x.*")
		var foundStack = false
		for _, line := range stackLines {
			if printFromHere == 0 {
				shortStack = append(shortStack, line)
			} else {
				if !foundStack && thisFuncEntryRE.MatchString(line) {
					foundStack = true
				}

				if foundStack {
					numBackedUp += 1
				}

				if numBackedUp >= printFromHere && numBackedUp < printFromHere+lengthOfEntry {
					shortStack = append(shortStack, line)
				} else if numBackedUp >= printFromHere+lengthOfEntry {
					break
				}
			}
		}
		return shortStack
	}
}

// numEntriesToSkip = the number of stack entries to "skip", meaning filter out.
//
//	We want to print only the line where the assertion is called.
//	So if you call this from an assertion, give 1.
//	If your code is one level above an assertion, give 2.
//	If you want to see the whole stack trace, give 0.
func (a *Asserter) PrintStack(numBackUp int) {
	var stackLines = a.GetRelevantStack(numBackUp + 1)
	for _, line := range stackLines {
		fmt.Println(line)
	}
}

func (a *Asserter) Equal(expected any, actual any) {
	if !(expected == actual) {
		a.PrintDivider()
		a.testingT.Errorf("Got %v (type %v), but expected to get %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
		a.PrintStack(1)
	}
}

func (a *Asserter) NotEqual(expected any, actual any) {
	if !(expected != actual) {
		a.PrintDivider()
		a.testingT.Errorf("values are equal and not expected to be: %v (type %v)", actual, reflect.TypeOf(actual))
		a.PrintStack(1)
	}
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr |
		~float32 | ~float64
}

func genericLT[NUM Number](a *Asserter, left NUM, right NUM) {
	if !(left < right) {
		a.PrintDivider()
		a.testingT.Errorf("left value (%v) is not less than the right value (%v)", left, right)
		a.PrintStack(1)
	}
}
func genericGT[NUM Number](a *Asserter, left NUM, right NUM) {
	if !(left > right) {
		a.PrintDivider()
		a.testingT.Errorf("left value (%v) is not greater than the right value (%v)", left, right)
		a.PrintStack(1)
	}
}
func genericLE[NUM Number](a *Asserter, left NUM, right NUM) {
	if !(left <= right) {
		a.PrintDivider()
		a.testingT.Errorf("left value (%v) is not less than or equal to the right value (%v)", left, right)
		a.PrintStack(1)
	}
}
func genericGE[NUM Number](a *Asserter, left NUM, right NUM) {
	if !(left >= right) {
		a.PrintDivider()
		a.testingT.Errorf("left value (%v) is not greater than or equal to the right value (%v)", left, right)
		a.PrintStack(2) // we are assuming this is called from an assert assertion function (defined below)
	}
}

// Phew... sure hope golang supports generic types in class methods sometime...
func (a *Asserter) LT(left any, right any) { // we assume these are the same type
	switch leftCasted := left.(type) {
	// , int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64,
	case int:
		switch rightCasted := right.(type) {
		case int:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int8:
		switch rightCasted := right.(type) {
		case int8:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int16:
		switch rightCasted := right.(type) {
		case int16:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int32:
		switch rightCasted := right.(type) {
		case int32:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int64:
		switch rightCasted := right.(type) {
		case int64:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint:
		switch rightCasted := right.(type) {
		case uint:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint16:
		switch rightCasted := right.(type) {
		case uint16:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint32:
		switch rightCasted := right.(type) {
		case uint32:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint64:
		switch rightCasted := right.(type) {
		case uint64:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uintptr:
		switch rightCasted := right.(type) {
		case uintptr:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float32:
		switch rightCasted := right.(type) {
		case float32:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float64:
		switch rightCasted := right.(type) {
		case float64:
			genericLT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	default:
		panic("cannot call this function with a non integral or non floating point type, or types that don't match (left)")
	}
}

func (a *Asserter) GT(left any, right any) { // we assume these are the same type
	switch leftCasted := left.(type) {
	// , int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64,
	case int:
		switch rightCasted := right.(type) {
		case int:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int8:
		switch rightCasted := right.(type) {
		case int8:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int16:
		switch rightCasted := right.(type) {
		case int16:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int32:
		switch rightCasted := right.(type) {
		case int32:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int64:
		switch rightCasted := right.(type) {
		case int64:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint:
		switch rightCasted := right.(type) {
		case uint:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint16:
		switch rightCasted := right.(type) {
		case uint16:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint32:
		switch rightCasted := right.(type) {
		case uint32:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint64:
		switch rightCasted := right.(type) {
		case uint64:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uintptr:
		switch rightCasted := right.(type) {
		case uintptr:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float32:
		switch rightCasted := right.(type) {
		case float32:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float64:
		switch rightCasted := right.(type) {
		case float64:
			genericGT(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	default:
		panic("cannot call this function with a non integral or non floating point type, or types that don't match (left)")
	}
}

func (a *Asserter) LE(left any, right any) { // we assume these are the same type
	switch leftCasted := left.(type) {
	// , int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64,
	case int:
		switch rightCasted := right.(type) {
		case int:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int8:
		switch rightCasted := right.(type) {
		case int8:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int16:
		switch rightCasted := right.(type) {
		case int16:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int32:
		switch rightCasted := right.(type) {
		case int32:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int64:
		switch rightCasted := right.(type) {
		case int64:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint:
		switch rightCasted := right.(type) {
		case uint:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint16:
		switch rightCasted := right.(type) {
		case uint16:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint32:
		switch rightCasted := right.(type) {
		case uint32:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint64:
		switch rightCasted := right.(type) {
		case uint64:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uintptr:
		switch rightCasted := right.(type) {
		case uintptr:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float32:
		switch rightCasted := right.(type) {
		case float32:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float64:
		switch rightCasted := right.(type) {
		case float64:
			genericLE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	default:
		panic("cannot call this function with a non integral or non floating point type, or types that don't match (left)")
	}
}

func (a *Asserter) GE(left any, right any) { // we assume these are the same type
	switch leftCasted := left.(type) {
	// , int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64,
	case int:
		switch rightCasted := right.(type) {
		case int:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int8:
		switch rightCasted := right.(type) {
		case int8:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int16:
		switch rightCasted := right.(type) {
		case int16:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int32:
		switch rightCasted := right.(type) {
		case int32:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case int64:
		switch rightCasted := right.(type) {
		case int64:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint:
		switch rightCasted := right.(type) {
		case uint:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint16:
		switch rightCasted := right.(type) {
		case uint16:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint32:
		switch rightCasted := right.(type) {
		case uint32:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uint64:
		switch rightCasted := right.(type) {
		case uint64:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case uintptr:
		switch rightCasted := right.(type) {
		case uintptr:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float32:
		switch rightCasted := right.(type) {
		case float32:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	case float64:
		switch rightCasted := right.(type) {
		case float64:
			genericGE(a, leftCasted, rightCasted)
		default:
			panic("cannot call this function with a non integral or non floating point type, or types that don't match (right)")
		}
	default:
		panic("cannot call this function with a non integral or non floating point type, or types that don't match (left)")
	}
}
