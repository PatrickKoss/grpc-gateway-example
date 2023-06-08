package domain

type User struct {
	Id      string `json:"id" validate:"required,min=1,max=63"`
	Name    string `json:"name" validate:"required,min=1,max=200,onlySmallLetters"`
	Email   string `json:"email" validate:"required,min=1,max=63,isValidEmail"`
	Surname string `json:"surname" validate:"required,min=1,max=200,onlySmallLetters"`
}
