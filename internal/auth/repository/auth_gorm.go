package repository

import (
	"SiamOutlet/internal/auth/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthGormRepository struct {
	db *gorm.DB
}

func NewAuthGormRepository(db *gorm.DB) *AuthGormRepository {
	return &AuthGormRepository{db: db}
}

// TODO : ระบบ Register Login ทั้งหมด
func (r *AuthGormRepository) Register(req *domain.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	return r.db.Create(&user).Error
}

// TODO : Login Login ทั้งหมด
func (r *AuthGormRepository) Login(req *domain.LoginRequest) (string, error) {
	var user domain.User
	if err := r.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", err
	}

	//TODO: Generate JWT Token
	jwtSecretKey := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["Role"] = user.Role
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
