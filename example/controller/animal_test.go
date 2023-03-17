package controller

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

type GetAnimalTestcase struct {
	Animals                []schemas.Animal
	TestAnimal             schemas.Animal
	ExpectedResponseStatus int
}

func (testcase GetAnimalTestcase) Setup(t *testing.T, s fibertest.Suite) error {
	suite := s.(*test.Suite)
	for _, animal := range testcase.Animals {
		suite.Data[animal.Name] = animal.Sound
	}
	return nil
}

func (GetAnimalTestcase) App(t *testing.T, s fibertest.Suite) (*fiber.App, error) {
	suite := s.(*test.Suite)
	testApp := Controller{Data: suite.Data}

	api := fiber.New()
	RegisterRoutes(api.Group("/api/v1"), testApp)

	return api, nil
}

func (testcase GetAnimalTestcase) GetRequest(t *testing.T) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, "/api/v1/animal", nil)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	q := url.Values{}
	q.Add("name", testcase.TestAnimal.Name)
	request.URL.RawQuery = q.Encode()

	return request, nil
}

func (testcase GetAnimalTestcase) AssertResponse(t *testing.T, _ fibertest.Suite, response *http.Response) bool {
	if !assert.Equal(t, testcase.ExpectedResponseStatus, response.StatusCode) {
		return false
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return false
	}

	switch response.StatusCode {
	case http.StatusOK:
		var animal schemas.Animal
		err = json.Unmarshal(body, &animal)
		if err != nil {
			t.Error(err)
			return false
		}

		if !assert.Equal(t, testcase.TestAnimal, animal) {
			return false
		}

	case http.StatusNotFound:
		if !assert.Equal(t, "I don't know what sound this animal makes ;-;", string(body)) {
			return false
		}

	default:
		t.Errorf("invalid response status code %d for animal sound API", response.StatusCode)
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

type AddAnimalTestcase struct {
	InitialAnimals []schemas.Animal
	TestAnimal     schemas.Animal
}

func (testcase AddAnimalTestcase) Setup(t *testing.T, s fibertest.Suite) error {
	suite := s.(*test.Suite)
	for _, animal := range testcase.InitialAnimals {
		suite.Data[animal.Name] = animal.Sound
	}
	return nil
}

func (AddAnimalTestcase) App(t *testing.T, s fibertest.Suite) (*fiber.App, error) {
	suite := s.(*test.Suite)
	testApp := Controller{Data: suite.Data}

	api := fiber.New()
	RegisterRoutes(api.Group("/api/v1"), testApp)

	return api, nil
}

func (testcase AddAnimalTestcase) GetRequest(t *testing.T) (*http.Request, error) {
	body, err := json.Marshal(testcase.TestAnimal)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "/api/v1/animal", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	q := url.Values{}
	q.Add("name", testcase.TestAnimal.Name)
	request.URL.RawQuery = q.Encode()

	return request, nil
}

func (testcase AddAnimalTestcase) AssertResponse(t *testing.T, s fibertest.Suite, response *http.Response) bool {
	if !assert.Equal(t, http.StatusCreated, response.StatusCode) {
		return false
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return false
	}

	var animal schemas.Animal
	err = json.Unmarshal(body, &animal)
	if err != nil {
		t.Error(err)
		return false
	}

	if !assert.Equal(t, testcase.TestAnimal, animal) {
		return false
	}

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
