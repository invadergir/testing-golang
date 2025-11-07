package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/suite"

	// chi:
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ChiTestSuite struct {
	suite.Suite
	Router *chi.Mux
	Server *httptest.Server
}

// Setup function for the entire Suite (like beforeAll())
func (suite *ChiTestSuite) SetupSuite() {
	// if testing a real service, we wouldn't implement the router functions here of course:
	// TODO make a real impl that we can test.
	suite.Router = setupRouter()
	suite.Server = httptest.NewServer(suite.Router)
	// go suite.startServer()
}

// For real tests this won't be necessary, we can call some other router maker with maybe some mocks:
func setupRouter() *chi.Mux {
	var r = chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	return r
}

// Teardown for the suite (like afterAll())
func (suite *ChiTestSuite) TearDownSuite() {
	suite.Server.Close()
}

// Test example - tests the /ping route returns 'pong':
// Notice no need to pass in the 't' testing context!  The suite
// class has assertions built-in which make that unnecessary.
func (s *ChiTestSuite) Test_Template() {
	var resp, err = http.Get(s.Server.URL + "/ping")

	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)

	bodyRaw, err := io.ReadAll(resp.Body)
	body := string(bodyRaw)
	s.Require().NoError(err)
	s.Equal("pong", body)
}

// This triggers the start of the test suite:
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ChiTestSuite))
}
