package api

import (
	"github.com/Josh2604/master-class/utils"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is supported or not
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
