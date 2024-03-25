package usecase

import (
	"context"
	"time"

	domain "github.com/ijul/be-monggo/domain/request"
	"github.com/ijul/be-monggo/internal/tokenutil"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUscase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

// Create implements domain.SignupUsecase.
func (s *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRepository.Create(ctx, user)
}

// CreateAccessToken implements domain.SignupUsecase.
func (s *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

// CreateRefreshToken implements domain.SignupUsecase.
func (s *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

// GetUserByEmail implements domain.SignupUsecase.
func (s *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRepository.GetByEmail(ctx, email)
}
