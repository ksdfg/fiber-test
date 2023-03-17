package schemas

type Animal struct {
	Name  string `json:"name"`
	Sound string `json:"sound"`
}

type AnimalSoundQuery struct {
	Name string `query:"name"`
}
