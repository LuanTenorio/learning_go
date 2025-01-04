package model

import (
	"errors"

	"github.com/LuanTenorio/learning_go/api/database"
)

type Personality struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

func FindPersonalities() []Personality {

	var personalities []Personality

	database.DB.Find(&personalities)

	return personalities
}

func FindPersonalityById(id int) (Personality, error) {
	var personality Personality
	var err error

	database.DB.First(&personality, id)

	if personality.Id == 0 {
		err = errors.New("Personality not faund")
	}

	return personality, err
}

func CreatePersonality(personality *Personality) error {
	database.DB.Create(&personality)
	if personality.Id == 0 {
		return errors.New("error when creating personality")
	}

	return nil
}

func DeletePersonality(id int) error {
	var personality Personality
	err := database.DB.Delete(&personality, id).Error

	if err != nil {
		return err
	}

	return nil
}
