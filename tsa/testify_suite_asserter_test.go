package tsa

// LEFTOFF: this is still basically the same as the normal asserter tests, minus name changes.
// TODO:  find a way to integrate into the test suite and maybe we can get that test context for free in the setup function.

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// // Asserter template test
// func Test_(t *testing.T) {
// 	var assert = MakeTestifySuiteAsserter(t)
// 	assert.True(true)
// }

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TestifySuiteAsserterTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Setup function
func (suite *TestifySuiteAsserterTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.

// Testify style 1
func (suite *TestifySuiteAsserterTestSuite) Test_SetupWorked1() {
	//fmt.Println("1 suite ptr = ", suite, "\n1 test context ptr = ", suite.T())
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	// assert.Equal(suite.T(), 66, suite.VariableThatShouldStartAtFive)
}

// Testify style 2
func (suite *TestifySuiteAsserterTestSuite) Test_SetupWorked2() {
	// fmt.Println("2 suite ptr = ", suite, "\n2 test context ptr = ", suite.T())
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	// assert.Equal(suite.T(), 66, suite.VariableThatShouldStartAtFive)
}

// Asserter style - doesn't work as expected when inside a test suite.
func (suite *TestifySuiteAsserterTestSuite) Test_SetupWorkedAlsoAsserterWorks() {
	var assert = MakeTestifySuiteAsserter(suite.T())
	assert.Equal(5, suite.VariableThatShouldStartAtFive)
	assert.NotEqual(6, suite.VariableThatShouldStartAtFive)
}

// Asserter style - doesn't work as expected when inside a test suite.
// TODO sometime try to use Asserter to wrap the fns inside assertions.go
func (suite *TestifySuiteAsserterTestSuite) Test_SetupWorkedAlsoAsserterWorks2() {
	var assert = MakeTestifySuiteAsserter(suite.T())
	assert.Equal(5, suite.VariableThatShouldStartAtFive)
	assert.NotEqual(6, suite.VariableThatShouldStartAtFive)
}

// This triggers the start of the test suite:
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestifySuiteAsserterTestSuite))
}
