package request

type CreatePersonRequest struct{
	Name string `validate:"required,min=3" json:"name"`
}