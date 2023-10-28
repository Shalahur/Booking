package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid return true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required check for required fields
func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (form *Form) Has(field string, request *http.Request) bool {
	fieldValue := request.Form.Get(field)

	if fieldValue == "" {
		return false
	}
	return true
}

// MinLength check string minimum length
func (form *Form) MinLength(field string, length int, request *http.Request) bool {

	fieldValue := request.Form.Get(field)
	if len(fieldValue) < length {
		form.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail check email is valid or not using asaskevich/govalidator lib
func (form *Form) IsEmail(field string) {

	fieldValue := form.Get(field)
	if !govalidator.IsEmail(fieldValue) {
		form.Errors.Add(field, "Invalid email address.")
	}
}
