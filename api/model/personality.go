package model

type Personality struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

func FindPersonalities() []Personality {
	personalities := []Personality{{"Luan", "dev asdasd asd asdasd"}, {"shakespirito", "puppet..."}}

	return personalities
}
