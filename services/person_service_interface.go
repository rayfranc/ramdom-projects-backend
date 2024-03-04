package services

import (
	"main/data/request"
	"main/data/response"
)

type PersonService interface{
	Create(person request.CreatePersonRequest)
	Update(person request.UpdatePersonRequest)
	Delete(personId int)
	FindById(personId int) response.PersonResponse
	FindAll() []response.PersonResponse
	Shuffle(shuffleRequest request.ShuffleRequest) (response.ShuffleResponse)
}


