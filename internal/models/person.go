package models

import "time"

type Gender int

const (
	MaleGender Gender = iota
	FemaleGender
)

func (g Gender) String() string {
	switch g {
	case MaleGender:
		return "male"
	case FemaleGender:
		return "female"
	}

	return ""
}

type Person struct {
	ID          int `gorm:"primaryKey;column:user_id"`
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      Gender
	Nationalize string
	createdAt   time.Time
}

func (Person) TableName() string { return "person" }
