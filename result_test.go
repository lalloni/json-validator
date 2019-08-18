package validator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	validator "github.com/lalloni/json-validator"
)

func TestValidationResult(t *testing.T) {

	tests := []struct {
		name   string
		errors []validator.ValidationError
		want   string
	}{
		{"no errors", nil, "no errors"},
		{"1 errors", []validator.ValidationError{{"(root)", "Bla bla"}}, "(root): Bla bla"},
		{"2 errors", []validator.ValidationError{{"(root)", "Bla bla"}, {"persona.id", "No puede ser"}}, "2 errors: (1) (root): Bla bla; (2) persona.id: No puede ser"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			r := &validator.ValidationResult{Errors: tt.errors}
			a.EqualValues(tt.want, r.String())
		})
	}
}
