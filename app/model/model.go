package model

type Book struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name" gorm:"unique;size:100;not null;"`
	GenreID int     `json:"genre_id" gorm:"not null"`
	Price   float64 `json:"price" gorm:"not null"`
	Amount  uint    `json:"amount" gorm:"not null"`
	Genre   Genre   `gorm:"foreignKey:GenreID"`
}

type Genre struct {
	ID   uint   `json:"id" gorm:"gorm:primaryKey"`
	Name string `json:"name" gorm:"unique;type:varchar(100);not null"`
}
