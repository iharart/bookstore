package model

type Book struct {
	ID      uint    `gorm:"primary_key"`
	Name    string  `json:"name" gorm:"unique;size:100;not null;"`
	GenreID int     `json:"genre_id" gorm:"not null"`
	Price   float64 `json:"price" gorm:"UNSIGNED NOT NULL; check:(price > 0)"`
	Amount  uint    `json:"amount" gorm:"not null;"`
	Genre   Genre   `gorm:"<-:false;foreignKey:GenreID"`
}

type Genre struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name" gorm:"unique;type:varchar(100);not null"`
}
