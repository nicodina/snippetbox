package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form represents form data
type Form struct{
	url.Values
	Errors errors
}

// NewForm retruns a newly created form object
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if required fields are empty, it adds an error in that case
func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// MaxLength checks if a field respects the lenght limit
func (form *Form) MaxLength(field string, limit int) {
	value := form.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > limit {
		form.Errors.Add(field, fmt.Sprintf("The field is too long (maximum is %d characters)", limit))
	}
}

// PermittedValues checks if values are valid
func (form *Form) PermittedValues(field string, opts ...string) {
	value := form.Get(field)
	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	form.Add(field, "This field is invalid")
}

// Valid returns true if there are no errors
func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}