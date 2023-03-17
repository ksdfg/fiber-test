package test

type Suite struct {
	Data map[string]string
}

func (suite *Suite) Setup() error {
	suite.Data = make(map[string]string)
	return nil
}

func (suite *Suite) Reset() error {
	for key := range suite.Data {
		delete(suite.Data, key)
	}
	return nil
}

func (*Suite) Teardown() error {
	return nil
}
