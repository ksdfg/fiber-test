package fibertest

import (
	"testing"
)

// RunTest will run a unit test in a reset test suite, and assert whether the API response is valid.
func RunTest(t *testing.T, suite Suite, testcase Testcase) bool {
	// Reset test suite to initialized state
	err := suite.Reset()
	if err != nil {
		return false
	}

	// Setup all data required for the unit test in the suite
	err = testcase.Setup(t, suite)
	if err != nil {
		return false
	}

	// Generate fiber app in order to test the API
	app, err := testcase.App(t, suite)
	if err != nil {
		return false
	}

	// Generate http request to make the API call to the endpoint being tested
	request, err := testcase.GetRequest(t)
	if err != nil {
		return false
	}

	// Test the endpoint using generated request
	response, err := app.Test(request)
	if err != nil {
		t.Error(err)
		return false
	}

	defer response.Body.Close()

	// Assert that the response is as expected
	return testcase.AssertResponse(t, suite, response)
}
