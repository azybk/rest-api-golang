package repository

import (
	"context"
	"database/sql"
	"rest-api-golang/domain"

	"github.com/doug-martin/goqu/v9"
)

type mediaRepository struct {
	db *goqu.Database
}

func NewMedia(con *sql.DB) domain.MediaRepository {
	return &mediaRepository{
		db: goqu.New("default", con),
	}
}

func(m mediaRepository) FindById(ctx context.Context, id string) (media domain.Media, err error) {
	// dataset := m.db.From("media").Where(goqu.C("id").Eq(id))
	dataset := m.db.From("media").Where(goqu.Ex{
		"id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &media)
	return
}

func (m mediaRepository) FindByIds(ctx context.Context, ids []string) (media []domain.Media, err error) {
	dataset := m.db.From("media").Where(goqu.C("id").In(ids))
	err = dataset.ScanStructsContext(ctx, &media)
	return
}

func (m mediaRepository) Save(ctx context.Context, media *domain.Media) error {
	executor := m.db.Insert("media").Rows(media).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}