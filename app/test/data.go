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
