package database

import (
	"errors"
	model "github.com/iharart/bookstore/app/model"
	"gorm.io/gorm"
)

func GetBookByID(id uint, db *gorm.DB) (model.Book, bool, error) {
	book := model.Book{}
	/*query := db.Select("books.*")
	query = query.Group("books.id")
	err := query.Where("books.id = ?", id).First(&book).Error*/
	err := db.First(&book, model.Book{ID: id}).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return book, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return book, false, nil
	}
	return book, true, nil
}

func GetAllBooks(db *gorm.DB) ([]model.Book, error) {
	books := []model.Book{}
	if err := db.Debug().Preload(model.GENRE).Order("name").Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func DeleteBook(id uint, db *gorm.DB) error {
	var book model.Book
	/*if err := db.Where("id = ? ", id).Delete(&book).Error; err != nil {
		return err
	}*/
	if err := db.Delete(&book, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, book *model.Book) error {
	if err := db.Save(&book).Error; err != nil {
		return err
	}
	return nil
}
