package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/unrealandychan/golang_api_full/util"
)

var validCurrency validator.Func = func (fieldLavel validator.FieldLevel)bool{
	if currency , ok := fieldLavel.Field().Interface().(string); ok{
		// check if currency is support or not
		return util.IsSupportedCurrency(currency)
	}
	return false
}
