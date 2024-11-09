package domain

import (
	"context"
	"database/sql"
	"rest-api-golang/dto"
)

const (
	JournalStatusInProgress  = "IN_PROGRESS"
	JournalStatusInCompleted = "COMPLETED"
)

type Journal struct {
	Id         string       `db:"id"`
	BookId     string       `db:"book_id"`
	StockCode  string       `db:"stock_code"`
	CustomerId string       `db:"customer_id"`
	Status     string       `db:"status"`
	BorrowedAt sql.NullTime `db:"borrowed_at"`
	ReturnedAt sql.NullTime `db:"returned_at"`
	DueAt      sql.NullTime `db:"due_at"`
}

type JournalSearch struct {
	CustomerId string
	Status     string
}

type JournalRepository interface {
	Find(ctx context.Context, se JournalSearch) ([]Journal, error)
	FindById(ctx context.Context, id string) (Journal, error)
	Save(ctx context.Context, journal *Journal) error
	Update(ctx context.Context, journal *Journal) error
}

type JournalService interface {
	Index(ctx context.Context, se JournalSearch) ([]dto.JournalData, error)
	Create(ctx context.Context, req dto.CreateJournalRequest) error
	Return(ctx context.Context, req dto.ReturnedJournalRequest) error
}
