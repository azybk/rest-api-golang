package dto

type BookData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Coverurl    string `json:"cover_url"`
	Description string `json:"description"`
	Isbn        string `json:"isbn"`
}

type BookStockData struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

type BookShowData struct {
	BookData
	Stocks []BookStockData `json:"stocks"`
}

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Isbn        string `json:"isbn" validate:"required"`
	CoverId     string `json:"cover_id"`
}

type UpdateBookRequest struct {
	Id          string `json:"-"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Isbn        string `json:"isbn" validate:"required"`
	CoverId     string `json:"cover_id"`
}
