package validation

import (
	"github.com/go-playground/validator/v10"
	"simple_bank/util"
)

var ValidCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if current, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(current)
	}
	return false
}
