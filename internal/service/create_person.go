package service

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/masonkmeyer/agify"
	"github.com/masonkmeyer/genderize"
	"github.com/masonkmeyer/nationalize"

	"effectiveMobileTest/internal/dto"
	"effectiveMobileTest/internal/models"
)

func (s *Service) CreatePerson(req *dto.CreatePersonRequest) (*dto.CreatePersonResponse, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Surname = strings.TrimSpace(req.Surname)
	req.Patronymic = strings.TrimSpace(req.Patronymic)

	if req.Name == "" || req.Surname == "" {
		return nil, fmt.Errorf("parse name, surname: %w", dto.ErrBadRequest)
	}
	person := models.Person{Name: req.Name, Surname: req.Surname}
	if req.Patronymic != "" {
		person.Patronymic = req.Patronymic
	}
	slog.Info("New person: ", slog.String("name", person.Name), slog.String("surname", person.Surname))

	age, err := GetAge(person.Name)
	if err != nil {
		return nil, fmt.Errorf("GetAge: %w", err)
	}
	person.Age = age
	slog.Info("Added", slog.Int("age", person.Age))

	gender, err := GetGender(person.Name)
	if err != nil {
		return nil, fmt.Errorf("GetGender: %w", err)
	}
	person.Gender = gender
	slog.Info("Added", slog.String("gender", person.Gender.String()))

	country, err := GetCountry(person.Name)
	if err != nil {
		return nil, fmt.Errorf("GetCountry: %w", err)
	}
	person.Nationalize = country
	slog.Info("Added", slog.String("nationalize", person.Nationalize))

	res := s.db.Create(&person)
	if res.Error != nil {
		return nil, fmt.Errorf("db create person: %w", res.Error)
	}
	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("zero users created")
	}
	return &dto.CreatePersonResponse{Ok: true}, nil
}

func GetCountry(name string) (string, error) {
	Client := nationalize.NewClient()
	prediction, _, err := Client.Predict(name)
	if err != nil {
		return "", err
	}

	countries := prediction.Country
	max := countries[0].Probability
	country := countries[0].CountryId
	for i := range countries {
		if max < countries[i].Probability {
			max = countries[i].Probability
			country = countries[i].CountryId
		}
	}
	return country, nil
}

func GetAge(name string) (int, error) {
	client := agify.NewClient()
	prediction, _, err := client.Predict(name)
	if err != nil {
		return 0, err
	}
	return prediction.Age, nil
}

func GetGender(name string) (models.Gender, error) {
	client := genderize.NewClient()
	prediction2, _, err := client.Predict(name)
	if err != nil {
		return 0, err
	}
	if prediction2.Gender == "male" {
		return models.MaleGender, nil
	} else {
		return models.FemaleGender, nil
	}
}
