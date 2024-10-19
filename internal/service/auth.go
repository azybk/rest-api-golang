package service

import (
	"context"
	"errors"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"rest-api-golang/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(conf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		conf:           conf,
		userRepository: userRepository,
	}
}

func (a authService) Login(ctx context.Context, request dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))

	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil

}
