package handler

import "effectiveMobileTest/internal/dto"

type service interface {
	CreatePerson(req *dto.CreatePersonRequest) (*dto.CreatePersonResponse, error)
	DeletePerson(req *dto.DeletePersonRequest) (*dto.DeletePersonResponse, error)
	UpdatePerson(req *dto.UpdatePersonRequest, id int) (*dto.UpdatePersonResponse, error)
	GetPersons(req *dto.GetPersonsRequest) (*dto.GetPersonsResponse, error)
}
