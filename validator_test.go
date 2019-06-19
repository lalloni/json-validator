package validator

import "testing"

func TestValidationResult_String(t *testing.T) {
	tests := []struct {
		name   string
		errors []ValidationError
		want   string
	}{
		{"no errors", nil, "0 errors"},
		{"1 errors", []ValidationError{{"(root)", "Bla bla"}}, "1 error: (root) Bla bla"},
		{"2 errors", []ValidationError{{"(root)", "Bla bla"}, {"persona.id", "No puede ser"}}, "2 errors: (root) Bla bla, persona.id No puede ser"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := &ValidationResult{
				Errors: tt.errors,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("ValidationResult.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
