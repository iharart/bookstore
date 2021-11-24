package handler

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Service struct {
	Router *mux.Router
	DB     *gorm.DB
}
