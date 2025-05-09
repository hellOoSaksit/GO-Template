package usecase

import "SiamOutlet/internal/auth/domain"

type AuthUsecase struct {
	repo domain.AuthRepository
}

func NewAuthUsecase(repo domain.AuthRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) Register(req *domain.RegisterRequest) error {
	return u.repo.Register(req)
}

func (u *AuthUsecase) Login(req *domain.LoginRequest) (string, error) {
	return u.repo.Login(req)
}
