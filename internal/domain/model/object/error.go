package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorUserNotFound,
}

const (
	ErrUser string = "user"
)

var (
	ErrorNone         error = nil
	ErrorUserNotFound error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + smodel.ErrNotFound)
)

func GetErrors() []error {
	return ERRORS
}
