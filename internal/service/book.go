package service

import (
	"context"
	"database/sql"
	"errors"
	"path"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"rest-api-golang/internal/config"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	cnf                 *config.Config
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	mediaRepository     domain.MediaRepository
}

func NewBook(cnf *config.Config,
	bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository,
	mediaRepository domain.MediaRepository) domain.BookService {
	return &bookService{
		cnf:                 cnf,
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
		mediaRepository:     mediaRepository,
	}
}

func (bs bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := bs.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	coverId := make([]string, 0)
	for _, v := range result {
		if v.CoverId.Valid {
			coverId = append(coverId, v.CoverId.String)
		}
	}

	covers := make(map[string]string)
	if len(coverId) > 0 {
		coversDb, _ := bs.mediaRepository.FindByIds(ctx, coverId)
		for _, v := range coversDb {
			covers[v.Id] = path.Join(bs.cnf.Server.Asset, v.Path)
		}
	}

	var data []dto.BookData
	for _, v := range result {
		var coverUrl string
		if v2, e := covers[v.CoverId.String]; e {
			coverUrl = v2
		}

		data = append(data, dto.BookData{
			Id:          v.Id,
			Title:       v.Title,
			Coverurl:    coverUrl,
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
			Code:   v.Code,
			Status: v.Status,
		})
	}

	var coverUrl string
	if result.CoverId.Valid {
		cover, _ := bs.mediaRepository.FindById(ctx, result.CoverId.String)
		if cover.Path != "" {
			coverUrl = path.Join(bs.cnf.Server.Asset, cover.Path)
		}
	}

	return dto.BookShowData{
		BookData: dto.BookData{
			Id:          result.Id,
			Title:       result.Title,
			Coverurl:    coverUrl,
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
		CoverId:     coverId,
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
