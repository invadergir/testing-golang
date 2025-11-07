package main

import (
	"testing"
)

// func TestNormalFail1(t *testing.T) {
// 	var actual = addNumbers(1, 2)
// 	var expected = 4
// 	if actual != expected {
// 		t.Errorf("got %d, expected %d", actual, expected)
// 	}
// }
// func TestNormalFail2(t *testing.T) {
// 	var actual = addNumbers(3, 4)
// 	var expected = 0
// 	if actual != expected {
// 		t.Errorf("got %d, expected %d", actual, expected)
// 	}
// }

/* 2 failing tests output (Normal / built-in test fw):
go test ./...
--- FAIL: TestNormalFail1 (0.00s)
    asserter_test.go:11: got 3, expected 4
--- FAIL: TestNormalFail2 (0.00s)
    asserter_test.go:18: got 7, expected 0
FAIL
FAIL	testing-go/cmd	0.003s
FAIL
make: *** [Makefile:17: test] Error 1
*/

func TestNormal(t *testing.T) {
	var actual = addNumbers(1, 2)
	var expected = 3
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestAsserter1(t *testing.T) {
	var assert = MakeAsserter(t)
	result := addNumbers(1, 2)
	assert.Equal(3, result)
	assert.NotEqual(4, result)
}

func TestAsserter2(t *testing.T) {
	var assert = MakeAsserter(t)
	result := addNumbers(2, 3)
	assert.Equal(5, result)
	assert.NotEqual(6, result)
}

func TestAsserterFail1(t *testing.T) {
	var runMe = true
	runMe = false // uncomment to skip this test
	if runMe {
		var assert = MakeAsserter(t)
		result := addNumbers(1, 2)
		//assert.Equal(11, result)
		assert.LT(6.9999, 7.0)
		assert.NotEqual(3, result)

	}
}

func TestAsserterFail2(t *testing.T) {
	var runMe = true
	runMe = false // uncomment to skip this test
	if runMe {
		var assert = MakeAsserter(t)
		result := addNumbers(2, 3)
		assert.Equal(22, result)
		assert.NotEqual(4, result)

	}
}

func TestGrowingSlice(t *testing.T) {
	var slice = make([]int, 0, 10)
	if len(slice) != 0 {
		t.Errorf("slice len should be 0")
	}
	if cap(slice) != 10 {
		t.Errorf("slice capacity should be 10")
	}

	slice = append(slice, 1)
	if len(slice) != 1 {
		t.Errorf("slice len should be 1")
	}
	if cap(slice) != 10 {
		t.Errorf("slice capacity should be 10")
	}
}

// func Test_GetRelevantStackWorks(t *testing.T) {
// 	var inputStack
// }
