package model

type Book struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `json:"name"`
	GenreID int     `json:"-"`
	Price   float64 `json:"price"`
	Amount  uint    `json:"amount"`
	Genre   Genre   `gorm:"foreignKey:GenreID"`
}

type Genre struct {
	ID   int    `gorm:"gorm:primaryKey" json:"id"`
	Name string `json:"name"`
}
