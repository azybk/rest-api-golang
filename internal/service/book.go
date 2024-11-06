package service

import (
	"context"
	"database/sql"
	"errors"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

func (bs bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := bs.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var data []dto.BookData
	for _, v := range result {
		data = append(data, dto.BookData{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Isbn:        v.Isbn,
		})
	}

	return data, nil
}

func (bs bookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {
	result, err := bs.bookRepository.FindById(ctx, id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	if result.Id == "" {
		return dto.BookShowData{}, domain.BookNotFound
	}

	stocks, err := bs.bookStockRepository.FindByBookId(ctx, result.Id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	stocksData := make([]dto.BookStockData, 0)
	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code: v.Code,
			Status: v.Status,
		})
	}

	return dto.BookShowData{
		BookData: dto.BookData {
			Id:          result.Id,
			Title:       result.Title,
			Description: result.Description,
			Isbn:        result.Isbn,
		},
		Stocks: stocksData,
	}, nil
}

func (bs bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	book := domain.Book{
		Id:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Isbn:        req.Isbn,
		CoverId: coverId,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}
	return bs.bookRepository.Save(ctx, &book)
}

func (bs bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := bs.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("book tidak ditemukan")
	}

	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.Isbn = req.Isbn
	persisted.CoverId = coverId
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return bs.bookRepository.Update(ctx, &persisted)
}

func (bs bookService) Delete(ctx context.Context, id string) error {
	persisted, err := bs.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("book tidak ditemukan")
	}

	err = bs.bookRepository.Delete(ctx, persisted.Id)
	if err != nil {
		return err
	}

	return bs.bookStockRepository.DeleteByBookId(ctx, persisted.Id)
}
