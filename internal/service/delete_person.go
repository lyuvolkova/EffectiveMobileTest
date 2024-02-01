package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"effectiveMobileTest/internal/dto"
	"effectiveMobileTest/internal/models"
)

func (s *Service) DeletePerson(req *dto.DeletePersonRequest) (*dto.DeletePersonResponse, error) {

	var person models.Person
	res := s.db.First(&person, req.ID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find person: %w", dto.ErrNotFound)
		}

		return nil, fmt.Errorf("find person: %w", res.Error)
	}

	res = s.db.Delete(&person)
	if res.Error != nil {
		return nil, fmt.Errorf("delete person: %w", res.Error)
	}

	return &dto.DeletePersonResponse{Ok: true}, nil
}
