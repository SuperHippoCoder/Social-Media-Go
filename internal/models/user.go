package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex,not null"`
	Password string `gorm:"not null"`
}

type UserModel struct {
	DB *gorm.DB
}

// Аунтефикация пользователя
func (u *UserModel) AuthenticateUser(username, password string) (*User, error) {
	var user User

	// Ищет пользователя с таким именем. Так как юзернейм уникален если такой уже есть, то выдаем ошибку
	if err := u.DB.Model(&User{}).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Invalid credentials")
		}

		return nil, err
	}

	// Хешируем пароль для безопасности
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid credentials")
	}

	// Возращаем созданного юзера
	return &user, nil
}

func (u *UserModel) RegisterUser(username, password string) error {
	var count int64
	if _ = u.DB.Model(&User{}).Where("username = ?", username).Count(&count); count >= 1 {
		return errors.New("User with this name already exsits")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Cannot hash password")
	}

	user := User{
		Username: username,
		Password: string(hashedPassword),
	}

	return u.DB.Create(&user).Error
}
