package services

import (
	"main/data/request"
	"main/data/response"
	model "main/models"
	"main/repository"
	helper "main/utils"
	"math/rand"
)

type PersonServiceImpl struct {
	PersonRepository repository.PersonRepository
	
}

// Create implements PersonService.
func (p *PersonServiceImpl) Create(person request.CreatePersonRequest) {
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

func (p *PersonServiceImpl) Shuffle(shuffle request.ShuffleRequest) response.ShuffleResponse{
	var res= []response.PersonTasks{}
	 var projects=shuffle.Projects
	for _, p := range shuffle.Persons {
		res=append(res ,response.PersonTasks{
			Name: p.Name,
			Projects: []response.Projects{},
		})
	}
	 for _,p:=range projects{
       per:=helper.Filter(res,func (t response.PersonTasks) bool{
		return len(t.Projects)==minors(res)
	   })
	   var rmd=rand.Intn(len(per))
	   per[rmd].Projects=append(per[rmd].Projects, response.Projects{Name: p.Name})
	   res[per[rmd].RealIndex]= response.PersonTasks{
		Name: per[rmd].Name,
		Projects: per[rmd].Projects,
	   }
	 }
	return response.ShuffleResponse{
		Persons: res,
	}
}

func NewPersonServiceImpl(personRepo repository.PersonRepository) PersonService {
	return &PersonServiceImpl{
		PersonRepository: personRepo,
	}
}

func minors(pers []response.PersonTasks) int {
	var m int
	for i, e := range pers {
     if i==0 || len(e.Projects) < m {
         m = len(e.Projects)
     }
 }
 return m
}