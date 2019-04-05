package validator

import (
	"github.com/lalloni/gojsonschema"
	"github.com/pkg/errors"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator.git/schemas"
)

type Validator interface {
	ValidatePersonaJSON(persona []byte) (*ValidationResult, error)
}

func New() (Validator, error) {
	schema, err := schemas.PersonaSchema()
	if err != nil {
		return nil, errors.Wrap(err, "getting persona schema")
	}
	return &validator{
		schema: schema,
	}, nil
}

type validator struct {
	schema *gojsonschema.Schema
}

func (v *validator) ValidatePersonaJSON(persona []byte) (*ValidationResult, error) {
	return ValidateJSON(v.schema, persona)
}

type ValidationResult struct {
	Errors []ValidationError
}

func (r *ValidationResult) Valid() bool {
	return len(r.Errors) == 0
}

type ValidationError struct {
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
}
