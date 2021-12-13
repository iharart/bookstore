package database

import (
	"errors"
	"github.com/iharart/bookstore/app/model"
	"gorm.io/gorm"
)

const (
	ErrorDbIsNil = "db is null"
)

func GetBookByID(id uint, db *gorm.DB) (model.Book, bool, error) {
	book := model.Book{}

	if err := DbCheck(db); err != nil {
		return book, false, err
	}
	err := db.Preload(model.GENRE).First(&book, model.Book{ID: id}).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return book, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return book, false, nil
	}
	return book, true, nil
}

func GetAllBooks(urlParams map[string]string, db *gorm.DB) ([]model.Book, error) {

	if err := DbCheck(db); err != nil {
		return nil, err
	}
	return FilterBooks(urlParams, db)
}

func DeleteBook(id uint, db *gorm.DB) error {
	var book model.Book

	if err := DbCheck(db); err != nil {
		return err
	}

	if err := db.Delete(&book, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(id uint, db *gorm.DB, book *model.Book) error {
	if err := DbCheck(db); err != nil {
		return err
	}
	if err := db.Preload(model.GENRE).Where("id = ?", id).Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func DbCheck(db *gorm.DB) error {
	if db == nil {
		return errors.New(ErrorDbIsNil)
	}
	return nil
}
