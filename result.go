package validator

import (
	"strconv"
	"strings"
)

type ValidationResult struct {
	Errors []ValidationError
}

func (r *ValidationResult) Valid() bool {
	return len(r.Errors) == 0
}

func (r *ValidationResult) String() string {
	l := len(r.Errors)
	sb := &strings.Builder{}
	sb.WriteString(strconv.Itoa(l))
	sb.WriteString(" error")
	if l != 1 {
		sb.WriteString("s")
	}
	for i, e := range r.Errors {
		if i == 0 {
			sb.WriteString(": ")
		}
		sb.WriteString(e.Field)
		sb.WriteString(" ")
		sb.WriteString(e.Description)
		if i < l-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

type ValidationError struct {
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
}
