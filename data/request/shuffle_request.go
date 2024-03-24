package request

type ShuffleRequest struct{
	Persons []CreatePersonRequest `json:"persons" binding:"required,min=3" `
	Projects[]CreatePersonRequest `json:"projects" binding:"required,min=3" `
}