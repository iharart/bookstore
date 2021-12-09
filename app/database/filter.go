package database

import (
	"github.com/iharart/bookstore/app/model"
	"github.com/iharart/bookstore/app/utils"
	"gorm.io/gorm"
)

func FilterBooks(urlParams map[string]string, db *gorm.DB) ([]model.Book, error) {
	var filterName string
	var filterGenreId string

	if param, ok := urlParams["name"]; ok {
		filterName = param
	}
	if param, ok := urlParams["genreId"]; ok {
		filterGenreId = param
	}

	books := []model.Book{}
	if (len(filterName) != 0) && (len(filterGenreId) != 0) {
		if err := db.Debug().Preload(model.GENRE).Where("name = ? AND genre_id = ? AND amount > ?", filterName,
			utils.StringToUint(filterGenreId), 0).Find(&books).Error; err != nil {
			return books, err
		}

	} else if len(filterName) != 0 {
		if err := db.Debug().Preload(model.GENRE).Where("name = ? AND amount > ?", filterName, 0).Find(&books).Error; err != nil {
			return books, err
		}
	} else if len(filterGenreId) != 0 {
		if err := db.Debug().Preload(model.GENRE).Where("genre_id = ? AND amount > ?", utils.StringToUint(filterGenreId), 0).Find(&books).Error; err != nil {
			return books, err
		}
	} else {
		if err := db.Debug().Preload(model.GENRE).Order("name").Find(&books, "amount > ?", 0).Error; err != nil {
			return books, err
		}
	}

	return books, nil
}
