package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct , embeds a url.Values struct
type Form struct {
	url.Values
	Errors errors
}

// New initialized a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required checks for required fields in a form post
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}

// MinLength checks for min length of a field
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d character long", length))
		return false
	}
	return true
}

// Valid returns true if there no errors associated with the form value or false if it has errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
