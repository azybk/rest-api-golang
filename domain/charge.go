package domain

import (
	"context"
	"database/sql"
)

type Charge struct {
	Id             string       `db:"id"`
	Journal_id     string       `db:"journal_id"`
	DaysLate       int          `db:"days_late"`
	Daily_late_fee int          `db:"daily_late_fee"`
	Total          int          `db:"total"`
	User_id        string       `db:"user_id"`
	CreatedAt      sql.NullTime `db:"created_at"`
}

type ChargeRepository interface {
	Save(ctx context.Context, charge *Charge) error
}
