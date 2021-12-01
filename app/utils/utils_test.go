package utils_test

import (
	"github.com/iharart/bookstore/app/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ErrorMessage = "The two numbers should be the same"
)

func TestStringToInt(t *testing.T) {
	actual := utils.StringToInt("-6")
	var expected int = -6
	assert.Equal(t, actual, expected, ErrorMessage)
}

func TestStringToUint(t *testing.T) {
	actual := utils.StringToUint("5")
	var expected uint = 5
	assert.Equal(t, actual, expected, ErrorMessage)
}
