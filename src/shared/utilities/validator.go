package utilities

import (
	"github.com/juanfer2/whorship_band/src/shared/utilities/validations"
)

func ValidatorStruct(s any) []error {
	return validations.ValidateStruct(s)
}
