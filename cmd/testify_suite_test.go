package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// // Asserter template test
// func Test_(t *testing.T) {
// 	var assert = MakeAsserter(t)
// 	assert.True(true)
// }

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ChiTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Setup function
func (suite *ChiTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.

// Testify style
func (suite *ChiTestSuite) Test_SetupWorked1() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// Testify style
func (suite *ChiTestSuite) Test_SetupWorked2() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// Asserter style - doesn't work as expected when inside a test suite.
func (suite *ChiTestSuite) Test_SetupWorkedAlsoAsserterWorks() {
	var assert = MakeAsserter(suite.T())
	assert.Equal(5, suite.VariableThatShouldStartAtFive)
	assert.NotEqual(6, suite.VariableThatShouldStartAtFive)
}

// Asserter style - doesn't work as expected when inside a test suite.
// TODO sometime try to use Asserter to wrap the fns inside assertions.go
func (suite *ChiTestSuite) Test_SetupWorkedAlsoAsserterWorks2() {
	var assert = MakeAsserter(suite.T())
	assert.Equal(5, suite.VariableThatShouldStartAtFive)
	assert.NotEqual(6, suite.VariableThatShouldStartAtFive)
}

// This triggers the start of the test suite:
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ChiTestSuite))
}
