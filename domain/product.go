package domain

type Product struct {
	Id          string  `json:"id" validate:"required,min=1,max=63"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Name        string  `json:"name" validate:"required,min=1,max=50"`
}
