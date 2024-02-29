package services

import (
	"main/data/request"
	"main/data/response"
	"main/helper"
	model "main/models"
	"main/repository"

	"github.com/go-playground/validator/v10"
)

type PersonServiceImpl struct {
	PersonRepository repository.PersonRepository
	validate         *validator.Validate
}

// Create implements PersonService.
func (p *PersonServiceImpl) Create(person request.CreatePersonRequest) {
	err:= p.validate.Struct(person)
	helper.ErrorPanic(err)
	personModel:=model.Person{
		Name: person.Name,
	}
	p.PersonRepository.Save(personModel)
}

// Delete implements PersonService.
func (p *PersonServiceImpl) Delete(personId int) {
	p.PersonRepository.Delete(personId)
}

// FindAll implements PersonService.
func (p *PersonServiceImpl) FindAll() []response.PersonResponse {
	result:=p.PersonRepository.FindAll()
	var persons []response.PersonResponse
	for _, value := range result {
		person:=response.PersonResponse{
			Id:value.Id,
			Name:value.Name,
		}
		persons = append(persons, person)
	}
	return persons
}

// FindById implements PersonService.
func (p *PersonServiceImpl) FindById(personId int) response.PersonResponse {
	personData,err:=p.PersonRepository.FindById(personId)
	helper.ErrorPanic(err)
	personResponse:= response.PersonResponse{
		Id: personData.Id,
		Name:  personData.Name,
	}
	return personResponse
}

// Update implements PersonService.
func (p *PersonServiceImpl) Update(person request.UpdatePersonRequest) {
	personData,err:=p.PersonRepository.FindById(person.Id)
	helper.ErrorPanic(err)
	personData.Name=person.Name
	p.PersonRepository.Update(personData)
}

func NewPersonServiceImpl(personRepo repository.PersonRepository, validate *validator.Validate) PersonService {
	return &PersonServiceImpl{
		PersonRepository: personRepo,
		validate: validate,
	}
}
