package test

import "github.com/iharart/bookstore/app/model"

var SampleGenresFail = []model.Genre{
	{
		Name: "Horror",
	},
	{
		Name: "Horror",
	},
}

var SampleGenresOk = []model.Genre{
	{
		Name: "Adventure",
	},
	{
		Name: "Classics",
	},
	{
		Name: "Fantasy",
	},
}

var SampleBooksAmountMore0 = []model.Book{
	{
		ID: 2,
		Name:    "Dracula",
		GenreID: 2,
		Price:   40.44,
		Amount:  30,
		Genre:   SampleGenreClassicsExpected,
	},
	{
		ID: 4,
		Name:    "Moby Dick",
		GenreID: 2,
		Price:   20.44,
		Amount:  10,
		Genre:   SampleGenreClassicsExpected,
	},
	{
		ID: 3,
		Name:    "The Three Musketeers",
		GenreID: 1,
		Price:   10.44,
		Amount:  5,
		Genre:   SampleGenreAdventureExpected,
	},
}

var SampleGetAllBooks = []model.Book{
	{
		Name:    "Game of thrones",
		GenreID: 3,
		Price:   30.44,
		Amount:  0,
		Genre:   SampleGenreFantasy,
	},
	{
		Name:    "Dracula",
		GenreID: 2,
		Price:   40.44,
		Amount:  30,
		Genre:   SampleGenreClassics,
	},
	{
		Name:    "The Three Musketeers",
		GenreID: 1,
		Price:   10.44,
		Amount:  5,
		Genre:   SampleGenreAdventure,
	},
	{
		Name:    "Moby Dick",
		GenreID: 2,
		Price:   20.44,
		Amount:  10,
		Genre:   SampleGenreClassics,
	},
}

var SampleBooksGenreIdSame = []model.Book{
	{
		ID : 2,
		Name:    "Dracula",
		GenreID: 2,
		Price:   40.44,
		Amount:  30,
		Genre:   SampleGenreClassicsExpected,
	},
	{
		ID : 4,
		Name:    "Moby Dick",
		GenreID: 2,
		Price:   20.44,
		Amount:  10,
		Genre:   SampleGenreClassicsExpected,
	},
}

var SampleBook = model.Book{
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   10.44,
	Amount:  5,
}

var SampleBookExpected = model.Book{
	ID: 1,
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   10.44,
	Amount:  5,
	Genre:   SampleGenreAdventureExpected,
}

var SampleBookNegativePrice = model.Book{
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   -100.44,
	Amount:  7,
}

var SampleGetBookById = model.Book{
	Name:    "The Great Gatsby",
	GenreID: 1,
	Price:   100,
	Amount:  5,
	Genre:   SampleGenreAdventureExpected,
}

var SampleUpdateBookOk = model.Book{
	Name:    "The Three Musketeers",
	GenreID: 1,
	Price:   19.44,
	Amount:  20,
	Genre:   SampleGenreAdventureExpected,
}

var SampleGenreAdventure = model.Genre{
	Name: "Adventure",
}

var SampleGenreClassics = model.Genre{
	Name: "Classics",
}

var SampleGenreFantasy = model.Genre{
	Name: "Fantasy",
}

var SampleGenreAdventureExpected = model.Genre{
	ID : 1,
	Name: "Adventure",
}

var SampleGenreClassicsExpected = model.Genre{
	ID : 2,
	Name: "Classics",
}

var SampleGenreFantasyExpected = model.Genre{
	ID : 3,
	Name: "Fantasy",
}
