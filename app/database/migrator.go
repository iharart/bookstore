package database

import (
	"github.com/iharart/bookstore/app/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) (*gorm.DB, error) {
	err := db.AutoMigrate(&model.Book{}, &model.Genre{})
	return db, err
}
