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
		Name: "Classics",
	},
	{
		ID:   3,
		Name: "Fantasy",
	},
}

var SampleBooksAmountMore0 = []model.Book{
	{
		ID:      4,
		Name:    "Dracula",
		GenreID: 2,
		Price:   40.44,
		Amount:  30,
		Genre:   SampleGenreClassics,
	},
	{
		ID:      6,
		Name:    "Moby Dick",
		GenreID: 2,
		Price:   20.44,
		Amount:  10,
		Genre:   SampleGenreClassics,
	},
	{
		ID:      5,
		Name:    "The Three Musketeers",
		GenreID: 1,
		Price:   10.44,
		Amount:  5,
		Genre:   SampleGenreAdventure,
	},
}

var SampleGetAllBooks = []model.Book{
	{
		ID:      3,
		Name:    "Game of thrones",
		GenreID: 3,
		Price:   30.44,
		Amount:  0,
		Genre:   SampleGenreFantasy,
	},
	{
		ID:      4,
		Name:    "Dracula",
		GenreID: 2,
		Price:   40.44,
		Amount:  30,
		Genre:   SampleGenreClassics,
	},
	{
		ID:      5,
		Name:    "The Three Musketeers",
		GenreID: 1,
		Price:   10.44,
		Amount:  5,
		Genre:   SampleGenreAdventure,
	},
	{
		ID:      6,
		Name:    "Moby Dick",
		GenreID: 2,
		Price:   20.44,
		Amount:  10,
		Genre:   SampleGenreClassics,
	},
}

var SampleBooksGenreIdSame = []model.Book{
	{
		ID:      4,
		Name:    "Dracula",
		GenreID: 2,
		Price:   40.44,
		Amount:  30,
		Genre:   SampleGenreClassics,
	},
	{
		ID:      6,
		Name:    "Moby Dick",
		GenreID: 2,
		Price:   20.44,
		Amount:  10,
		Genre:   SampleGenreClassics,
	},
}

var SampleBook = model.Book{
	ID:      5,
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   10.44,
	Amount:  5,
	Genre:   SampleGenreAdventure,
}

var SampleBookNegativePrice = model.Book{
	ID:      5,
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   -100.44,
	Amount:  7,
}

var SampleGetBookById = model.Book{
	ID:      6,
	Name:    "The Great Gatsby",
	GenreID: 1,
	Price:   100,
	Amount:  5,
	Genre:   SampleGenreAdventure,
}

var SampleUpdateBookOk = model.Book{
	ID:      5,
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   19.44,
	Amount:  20,
	Genre:   SampleGenreAdventure,
}

var SampleGenreAdventure = model.Genre{
	ID:   1,
	Name: "Adventure",
}

var SampleGenreClassics = model.Genre{
	ID:   2,
	Name: "Classics",
}

var SampleGenreFantasy = model.Genre{
	ID:   1,
	Name: "Fantasy",
}
