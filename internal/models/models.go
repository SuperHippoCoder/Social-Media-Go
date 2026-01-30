package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBModel struct {
	User UserModel
	DB   *gorm.DB
}

// Инициализация датабазы
func InitDB(dataSourceName string) (*DBModel, error) {
	// Открываем датабазу
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %v", err)
	}

	// Добавляем таблицу в неё
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}

	//Создаем модель дб
	dbModel := &DBModel{
		DB:   db,
		User: UserModel{DB: db},
	}

	// Возращаем созданную
	return dbModel, nil
}
