package validation

import (
	"simple_bank/util"

	"github.com/go-playground/validator/v10"
)

var ValidCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if current, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(current)
	}
	return false
}

var BankValidatior = validator.New()

func Init() {
	validator := validator.New()
	// register custom validation
	BankValidatior = validator
}
