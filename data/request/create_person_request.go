package request

type CreatePersonRequest struct{
	Name string `json:"name" binding:"required,min=3" `
}