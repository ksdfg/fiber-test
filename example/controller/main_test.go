package controller

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

	suite = &test.Suite{}

	err := suite.Setup()
	if err != nil {
		panic(err)
	}

	code := m.Run()

	err = suite.Teardown()
	if err != nil {
		panic(err)
	}

	os.Exit(code)
}
