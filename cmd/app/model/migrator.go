package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{}, &Genre{})
	return db
}
