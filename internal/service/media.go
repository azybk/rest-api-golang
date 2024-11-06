package service

import (
	"context"
	"database/sql"
	"path"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"rest-api-golang/internal/config"
	"time"

	"github.com/google/uuid"
)

type mediaService struct {
	conf *config.Config
	mediaRepository domain.MediaRepository
}

func NewMedia(conf *config.Config, mediaRepository domain.MediaRepository) domain.MediaService {
	return &mediaService{
		conf: conf,
		mediaRepository: mediaRepository,
	}
}

func (m mediaService) Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error) {
	media := domain.Media{
		Id: uuid.NewString(),
		Path: req.Path,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	err := m.mediaRepository.Save(ctx, &media)
	if err != nil {
		return dto.MediaData{}, err
	}

	url := path.Join(m.conf.Server.Asset, media.Path)
	return dto.MediaData{
		Id: media.Id,
		Path: media.Path,
		Url: url,
	}, nil
}