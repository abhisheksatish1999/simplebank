package api

import (
	"github.com/abhisheksatish1999/simplebank/util"
	"github.com/go-playground/validator/v10"
)

// Validator function for currency field
var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
