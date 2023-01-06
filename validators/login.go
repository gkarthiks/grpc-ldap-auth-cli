package validators

import (
	"errors"
)

var UserNameValidate = func(input string) error {
	if len(input) <= 3 {
		return errors.New("please provide a valid username")
	}
	return nil
}

var PasswdLenValidate = func(input string) error {
	if len(input) < 1 {
		return errors.New("please provide a password")
	}
	return nil
}
