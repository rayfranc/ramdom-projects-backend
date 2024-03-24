package repository

import (
	model "main/models"
)

type PersonRepository interface {
	Save(tags model.Person)
	Update(tags model.Person)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Person, err error)
	FindAll() []model.Person
}