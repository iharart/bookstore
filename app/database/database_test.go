package database_test

import (
	"github.com/iharart/bookstore/app/database"
	"github.com/iharart/bookstore/app/model"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

var DB *gorm.DB

func TestDbCheck(t *testing.T) {
	err := database.DbCheck(DB)
	require.Error(t, err)
}

func TestDeleteBook(t *testing.T) {
	err := database.DeleteBook(1, DB)
	require.Error(t, err)
}

func TestGetGetAllBooks(t *testing.T) {
	_, err := database.GetAllBooks(DB)
	require.Error(t, err)
}

func TestUpdateBook(t *testing.T) {
	book := model.Book{}
	err := database.UpdateBook(DB, &book)
	require.Error(t, err)
}

func TestGetBookByID(t *testing.T) {
	_, _, err := database.GetBookByID(1, DB)
	require.Error(t, err)
}
