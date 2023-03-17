# fiber-test

A small framework that helps you set up and run unit tests for your gofiber app.

## How to use

### TestMain func

```go
package foo

import (
	"os"
	"testing"

	fibertest "github.com/ksdfg/fiber-test"

	"github.com/user/repo/test"
)

var suite fibertest.TestSuite

func TestMain(m *testing.M) {
	suite = &test.Suite{}

	// Setup test suite
	err := suite.Setup()
	if err != nil {
		panic(err)
	}

	// Run tests
	code := m.Run()

	// Teardown test suite
	err = suite.Teardown()
	if err != nil {
		panic(err)
	}

	os.Exit(code)
}

```

### TestFeature func

```go
package foo

import (
	"testing"

	fibertest "github.com/ksdfg/fiber-test"

	"github.com/user/repo/test"
)

func TestFeature(t *testing.T) {
	// Initialize test case
	var testcase fibertest.TestCase
	testcase = test.TestCase{}

	// Run test
	fibertest.RunTest(t, suite, testcase)
}

```
