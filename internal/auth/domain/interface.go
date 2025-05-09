package domain

type AuthRepository interface {
	Register(user *RegisterRequest) error
	Login(user *LoginRequest) (string, error)
}
