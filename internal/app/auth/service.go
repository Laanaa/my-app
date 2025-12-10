package auth

import (
	"errors"

	"github.com/Laanaa/my-app/internal/app/user"
	"golang.org/x/crypto/bcrypt"
)


type AuthService struct{
	Repo user.UserRepository
}

func (s AuthService) Register(name, email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &user.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}
	return s.Repo.Create(user)
}

func (s AuthService) Login(email, password string) (user.User, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return	user, errors.New("invalid email or password")
	}

	return user, nil
}
