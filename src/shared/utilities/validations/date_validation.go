package validations

import (
	"fmt"
	"regexp"
)

var dateRe = regexp.MustCompile(`^(?:19|20)\d\d-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12][0-9]|3[01])$`)

type DateValidator struct{}

func (v DateValidator) Validate(val any) (bool, error) {
	if len(val.(string)) == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if !dateRe.MatchString(val.(string)) {
		return false, fmt.Errorf("is not correct format - (YYYY-MM-DD)")
	}
	return true, nil
}
