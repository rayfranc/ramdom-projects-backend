package request

type UpdatePersonRequest struct{
	Id int `validate:"required"`
	Name string `validate:"required,min=3" json:"name"`
}