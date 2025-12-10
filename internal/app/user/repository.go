package user

import "github.com/Laanaa/my-app/internal/database"

type UserRepository struct{}

func (r UserRepository) FindByEmail(email string) (User, error) {
	var user User
	err := database.DB.Where("email= ?", email).First(&user).Error
	return user, err
}

func (r UserRepository) Create(user *User) error {
	return database.DB.Create(user).Error
}