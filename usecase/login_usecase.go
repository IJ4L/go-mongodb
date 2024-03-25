package usecase

import (
	"context"
	"time"

	domain "github.com/ijul/be-monggo/domain/request"
	"github.com/ijul/be-monggo/internal/tokenutil"
)

type loginUsecase struct {
	userepository  domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userepository:  userRepository,
		contextTimeout: timeout,
	}
}

// CreateAccessToken implements domain.LoginUsecase.
func (l *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

// CreateRefreshToken implements domain.LoginUsecase.
func (l *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

// GetUserByEmail implements domain.LoginUsecase.
func (l *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.userepository.GetByEmail(ctx, email)
}
