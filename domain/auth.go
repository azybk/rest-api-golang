package domain

import (
	"context"
	"rest-api-golang/dto"
)

type AuthService interface {
	Login(ctx context.Context, request dto.AuthRequest) (dto.AuthResponse, error)
}
