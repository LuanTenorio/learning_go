package model

import (
	"errors"
)

type Personality struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

func FindPersonalities() []Personality {
	personalities := []Personality{
		{1, "Luan", "dev asdasd asd asdasd"},
		{2, "shakespirito", "puppet..."}}

	return personalities
}

func FindPersonalityById(id int) (Personality, error) {
	for _, personality := range FindPersonalities() {
		if personality.Id == id {
			return personality, nil
		}
	}

	return Personality{}, errors.New("Personality not faund")
}
