package validator_test

import (
	"testing"

	validator "gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator.git"
)

func TestValidationResult(t *testing.T) {
	tests := []struct {
		name   string
		errors []validator.ValidationError
		want   string
	}{
		{"no errors", nil, "0 errors"},
		{"1 errors", []validator.ValidationError{{"(root)", "Bla bla"}}, "1 error: (root) Bla bla"},
		{"2 errors", []validator.ValidationError{{"(root)", "Bla bla"}, {"persona.id", "No puede ser"}}, "2 errors: (root) Bla bla, persona.id No puede ser"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := &validator.ValidationResult{
				Errors: tt.errors,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("ValidationResult.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
