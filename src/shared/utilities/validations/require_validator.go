package validations

import (
	"fmt"
)

type RequireValidator struct {
	Expression string
}

func (v RequireValidator) Validate(val any) (bool, error) {
	l := len(val.(string))
	if l == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	return true, nil
}
