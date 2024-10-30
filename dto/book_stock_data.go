package dto

type CreateBookStockRequest struct {
	BookId string   `json:"book_id" validate:"required"`
	Codes  []string `json:"code" validate:"required,unique,min=1"`
}

type DeleteBookStockRequest struct {
	Codes []string
}