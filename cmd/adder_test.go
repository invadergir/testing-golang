package main

import (
	"reflect"
	"testing"
)

// type tester struct {
// 	tt *testing.T
// }
// func (t *tester) equals[C comparable](actual C, expected C) bool {
// 	fmt.Print("")
// 	if !(actual == expected) {
// 		t.tt.Error("expected "+expected+" but instead got ", actual)
// 	}
// }

func assertEqual[C comparable](t *testing.T, expected C, actual C) {
	if !(expected == actual) {
		t.Errorf("Got %v (type %v), but expected to get %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}
func assertNotEqual[C comparable](t *testing.T, expected C, actual C) {
	if !(expected != actual) {
		t.Errorf("values are equal and not expected to be: %v (type %v)", actual, reflect.TypeOf(actual))
	}
}

// func Test_addNumbers(t *testing.T) {
// 	result := addNumbers(2, 3)
// 	assertEqual(t, 5, result)
// 	assertNotEqual(t, 6, result)
// 	// if result != 5 {
// 	// 	t.Error("expected 5, got  ", result)
// 	// }

// }
