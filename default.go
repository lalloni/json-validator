package validator

import "fmt"

var Default = MustNew()

func MustNew() Validator {
	v, err := New()
	if err != nil {
		panic(fmt.Sprintf("Error creating new validator: %v", err))
	}
	return v
}
