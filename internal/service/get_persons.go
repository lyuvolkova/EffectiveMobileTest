package service

import (
	"fmt"
	"strings"

	"effectiveMobileTest/internal/dto"
	"effectiveMobileTest/internal/models"
)

const limit = 5

func (s *Service) GetPersons(req *dto.GetPersonsRequest) (*dto.GetPersonsResponse, error) {
	var persons []models.Person
	var person models.Person

	if req.Name = strings.TrimSpace(req.Name); req.Name != "" {
		person.Name = req.Name
	}
	if req.Surname = strings.TrimSpace(req.Surname); req.Surname != "" {
		person.Surname = req.Surname
	}
	if req.Patronymic = strings.TrimSpace(req.Patronymic); req.Patronymic != "" {
		person.Patronymic = req.Patronymic
	}

	offset := (req.Page - 1) * limit

	var total int64
	res := s.db.Model(person).Where(&person).Count(&total)
	if res.Error != nil {
		return nil, fmt.Errorf("get total count: %w", res.Error)
	}

	res = s.db.Limit(limit).Offset(offset).Where(&person).Find(&persons)
	if res.Error != nil {
		return nil, fmt.Errorf("getPerson: %w", res.Error)
	}
	personsDTO := make([]dto.Person, len(persons))
	for i := range persons {
		personsDTO[i] = dto.Person{
			ID:          persons[i].ID,
			Name:        persons[i].Name,
			Surname:     persons[i].Surname,
			Patronymic:  persons[i].Patronymic,
			Age:         persons[i].Age,
			Gender:      persons[i].Gender.String(),
			Nationalize: persons[i].Nationalize,
		}
	}

	return &dto.GetPersonsResponse{Limit: limit, TotalCount: total, Persons: personsDTO}, nil
}
