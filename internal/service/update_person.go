package service

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"effectiveMobileTest/internal/dto"
	"effectiveMobileTest/internal/models"
)

func (s *Service) UpdatePerson(req *dto.UpdatePersonRequest, id int) (*dto.UpdatePersonResponse, error) {
	var person models.Person
	var err error
	person.ID = id
	res := s.db.First(&person)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("person: %w", dto.ErrNotFound)
		}

		return nil, fmt.Errorf("get person: %w", res.Error)
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name != "" {
		person.Name = req.Name
		person.Age, err = GetAge(req.Name)
		if err != nil {
			return nil, fmt.Errorf("GetAge: %w", err)
		}
		person.Gender, err = GetGender(req.Name)
		if err != nil {
			return nil, fmt.Errorf("GetGender: %w", err)
		}
		person.Nationalize, err = GetCountry(req.Name)
		if err != nil {
			return nil, fmt.Errorf("GetCountry: %w", err)
		}
	}
	req.Surname = strings.TrimSpace(req.Surname)
	req.Patronymic = strings.TrimSpace(req.Patronymic)

	if req.Surname != "" {
		person.Surname = req.Surname
	}
	if req.Patronymic != "" {
		person.Patronymic = req.Patronymic
	}

	res = s.db.Save(&person)
	if res.Error != nil {
		return nil, fmt.Errorf("save person: %w", err)
	}

	return &dto.UpdatePersonResponse{Ok: true}, nil
}
