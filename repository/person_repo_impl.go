package repository

import (
	"fmt"
	"main/data/request"
	model "main/models"
	helper "main/utils"

	"gorm.io/gorm"
)

type PersonRepositoryImpl struct {
	Db *gorm.DB
}

func NewPersonRepositoryImpl(Db *gorm.DB) PersonRepository {
	return &PersonRepositoryImpl{Db: Db}
}

func (t PersonRepositoryImpl) Save(person model.Person) {
	result := t.Db.Create(&person)
	helper.ErrorPanic(result.Error)

}

func (t PersonRepositoryImpl) Update(person model.Person) {
	var updatePerson = request.UpdatePersonRequest{
		Id:   person.Id,
		Name: person.Name,
	}
	result := t.Db.Model(&person).Updates(updatePerson)
	helper.ErrorPanic(result.Error)
}

func (t PersonRepositoryImpl) Delete(personId int) {
	var person model.Person
	result := t.Db.Where("id = ?", personId).Delete(&person)
	helper.ErrorPanic(result.Error)
}

func (t PersonRepositoryImpl) FindById(personId int) (model.Person, error) {
	var person model.Person
	result := t.Db.Where("id = ?", personId).First(&person)
	if result.Error == nil {
		fmt.Println(result.Error)
		return person, nil
	} else {
		return person, result.Error
		
	}
}

func (t PersonRepositoryImpl) FindAll() []model.Person {
	var person []model.Person
	results := t.Db.Find(&person)
	helper.ErrorPanic(results.Error)
	return person
}