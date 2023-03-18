package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	fibertest "github.com/ksdfg/fiber-test"

	"github.com/ksdfg/fiber-test/example/schemas"
	"github.com/ksdfg/fiber-test/example/test"
)

// GetAnimalTestcase consolidates all required variables and methods for running a unit test for the GetAnimal endpoint
type GetAnimalTestcase struct {
	Animals                []schemas.Animal
	TestAnimal             schemas.Animal
	ExpectedResponseStatus int
}

func (testcase GetAnimalTestcase) Setup(t *testing.T, s fibertest.Suite) error {
	// Add all the initial animals in the suite dataset
	suite := s.(*test.Suite)
	for _, animal := range testcase.Animals {
		suite.Data[animal.Name] = animal.Sound
	}
	return nil
}

func (GetAnimalTestcase) App(t *testing.T, s fibertest.Suite) (*fiber.App, error) {
	// Initialize api with the suite dataset
	suite := s.(*test.Suite)
	return GenApp(Controller{Data: suite.Data}), nil
}

func (testcase GetAnimalTestcase) GetRequest(t *testing.T) (*http.Request, error) {
	// Generate GET request to /api/v1/animal endpoint
	request, err := http.NewRequest(http.MethodGet, "/api/v1/animal", nil)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	// Add test animal name as a query param
	q := url.Values{}
	q.Add("name", testcase.TestAnimal.Name)
	request.URL.RawQuery = q.Encode()

	return request, nil
}

func (testcase GetAnimalTestcase) AssertResponse(t *testing.T, _ fibertest.Suite, response *http.Response) bool {
	// Assert response status code matches the expected response
	if !assert.Equal(t, testcase.ExpectedResponseStatus, response.StatusCode) {
		return false
	}

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return false
	}

	switch response.StatusCode {
	case http.StatusOK:
		// Parse response body to an animal
		var animal schemas.Animal
		err = json.Unmarshal(body, &animal)
		if err != nil {
			t.Error(err)
			return false
		}

		// Assert that the response animal matches the test animal
		if !assert.Equal(t, testcase.TestAnimal, animal) {
			return false
		}

	case http.StatusNotFound:
		// Assert that the error response text matches the expected response
		if !assert.Equal(t, "I don't know what sound this animal makes ;-;", string(body)) {
			return false
		}

	default:
		// This should never be hit, unless you configured your test case wrong like a dum dum
		t.Errorf("invalid response status code %d for animal sound API", response.StatusCode)
		return false
	}

	return true
}

func TestController_GetAnimal_Found(t *testing.T) {
	testcase := GetAnimalTestcase{
		Animals:                []schemas.Animal{{Name: "cat", Sound: "meow"}, {Name: "dog", Sound: "woof"}},
		TestAnimal:             schemas.Animal{Name: "cat", Sound: "meow"},
		ExpectedResponseStatus: http.StatusOK,
	}
	fibertest.RunTest(t, suite, testcase)
}

func TestController_GetAnimal_NotFound(t *testing.T) {
	testcase := GetAnimalTestcase{
		Animals:                []schemas.Animal{{Name: "cow", Sound: "moo"}, {Name: "dog", Sound: "woof"}},
		TestAnimal:             schemas.Animal{Name: "cat", Sound: "meow"},
		ExpectedResponseStatus: http.StatusNotFound,
	}
	fibertest.RunTest(t, suite, testcase)
}

// GetAnimalTestcase consolidates all required variables and methods for running a unit test for the AddAnimal endpoint
type AddAnimalTestcase struct {
	InitialAnimals []schemas.Animal
	TestAnimal     schemas.Animal
}

func (testcase AddAnimalTestcase) Setup(t *testing.T, s fibertest.Suite) error {
	// Add all the initial animals in the suite dataset
	suite := s.(*test.Suite)
	for _, animal := range testcase.InitialAnimals {
		suite.Data[animal.Name] = animal.Sound
	}
	return nil
}

func (AddAnimalTestcase) App(t *testing.T, s fibertest.Suite) (*fiber.App, error) {
	// Initialize api with the suite dataset
	suite := s.(*test.Suite)
	return GenApp(Controller{Data: suite.Data}), nil
}

func (testcase AddAnimalTestcase) GetRequest(t *testing.T) (*http.Request, error) {
	// Marshal test animal to json []bye body
	body, err := json.Marshal(testcase.TestAnimal)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	// Generate POST request to /api/v1/animal endpoint with body generated above
	request, err := http.NewRequest(http.MethodPost, "/api/v1/animal", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
		return nil, err
	}

	// Add content type header so that the controller can parse the body
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

func (testcase AddAnimalTestcase) AssertResponse(t *testing.T, s fibertest.Suite, response *http.Response) bool {
	// Assert response status code matches the expected response
	// Since 201 is the only valid response, this does not need to be configured in the Testcase object
	if !assert.Equal(t, http.StatusCreated, response.StatusCode) {
		return false
	}

	// Since the API adds the animal to the dataset,
	// check the suite dataset if the animal's data has been added correctly
	suite := s.(*test.Suite)
	if sound, ok := suite.Data[testcase.TestAnimal.Name]; !ok {
		t.Error("new animal not present in data")
		return false
	} else if sound != testcase.TestAnimal.Sound {
		t.Errorf("new animal added with the wrong sound; expected=%s, actual=%s", testcase.TestAnimal.Sound, sound)
		return false
	}

	return true
}

func TestController_AddAnimal(t *testing.T) {
	testcase := AddAnimalTestcase{
		InitialAnimals: []schemas.Animal{{Name: "cow", Sound: "moo"}, {Name: "dog", Sound: "woof"}},
		TestAnimal:     schemas.Animal{Name: "cat", Sound: "meow"},
	}
	fibertest.RunTest(t, suite, testcase)
}
