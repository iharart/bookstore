package test

import "github.com/iharart/bookstore/app/model"

var SampleGenresFail = []model.Genre{
	{
		ID:   4,
		Name: "Horror",
	},
	{
		ID:   5,
		Name: "Horror",
	},
}

var SampleGenresOk = []model.Genre{
	{
		ID:   1,
		Name: "Adventure",
	},
	{
		ID:   2,
		Name: "Classic",
	},
	{
		ID:   3,
		Name: "Fantasy",
	},
}

var SampleBook = model.Book{
	ID:      5,
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   10.44,
	Amount:  5,
}

var SampleGetBookById = model.Book{
	ID:      6,
	Name:    "The Great Gatsby",
	GenreID: 1,
	Price:   100,
	Amount:  5,
}

var SampleUpdateBookOk = model.Book{
	ID:      5,
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   19.44,
	Amount:  20,
}
