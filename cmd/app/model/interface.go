package model

const (
	BOOK  string = "Book"
	GENRE string = "Genre"
)

type Tabler interface {
	TableName() string
}

func (Book) TableName() string {
	return BOOK
}

func (Genre) TableName() string {
	return GENRE
}
