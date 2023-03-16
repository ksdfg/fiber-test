package fibertest

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// Suite represents a global suite that is responsible for either mocking or managing any dependencies like
// databases, redis caches, rabbitmq queues etc.
type Suite interface {
	// Setup is responsible for spinning up and initializing all dependencies.
	Setup() error

	// Reset is responsible for returning all dependencies to their initialized state.
	// This includes deleting all rows created in database tables, flushing all keys stored in redis cache etc.
	//
	// This allows all tests to be run on a clear slate, and running one test cannot affect the outcome of the next
	// test.
	Reset() error

	// Teardown is responsible for stopping and removing all the dependencies.
	Teardown() error
}

// Testcase represents a single unit test that needs to be run.
// The test case object consolidates all business logic required to actually run a unit test,
// in order to provide structure to test routines.
type Testcase interface {
	// Setup is responsible for setting up all the dependencies for the unit test.
	// This includes inserting data into the DB, adding keys into the cache etc.
	Setup(*testing.T, Suite) error

	// App returns a *fiber.App object that has the API endpoint that needs to be tested configured to use mocks or
	// suite resources for any dependencies.
	App(*testing.T, Suite) (*fiber.App, error)

	// GetRequest will generate the *http.Request object that can be used to perform the unit test by calling the API.
	GetRequest(*testing.T) (*http.Request, error)

	// AssertResponse will assert that the *http.Response (generated as a result of making the request generated by
	// GetRequest) has the expected status, headers and body.
	AssertResponse(*testing.T, Suite, *http.Response) bool
}