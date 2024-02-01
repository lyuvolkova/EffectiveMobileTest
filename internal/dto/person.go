package dto

type CreatePersonRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type CreatePersonResponse struct {
	Ok bool `json:"ok"`
}

type DeletePersonRequest struct {
	ID int `json:"userID"`
}

type DeletePersonResponse struct {
	Ok bool `json:"ok"`
}

type UpdatePersonRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UpdatePersonResponse struct {
	Ok bool `json:"ok"`
}

type GetPersonsRequest struct {
	Name       string `json:"-"`
	Surname    string `json:"-"`
	Patronymic string `json:"-"`
	Page       int    `json:"-"`
}

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationalize string `json:"nationalize"`
}

type GetPersonsResponse struct {
	Limit      int      `json:"limit"`
	TotalCount int64    `json:"total_count"`
	Persons    []Person `json:"persons"`
}
