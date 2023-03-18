package api

import (
	"log"
	"os"
	"testing"

	fibertest "github.com/ksdfg/fiber-test"

	"github.com/ksdfg/fiber-test/example/test"
)

var suite fibertest.Suite

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialize test suite
	suite = &test.Suite{}
	err := suite.Setup()
	if err != nil {
		panic(err)
	}

	// Run all tests
	code := m.Run()

	// Teardown test suite
	err = suite.Teardown()
	if err != nil {
		panic(err)
	}

	os.Exit(code)
}
