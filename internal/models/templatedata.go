package models

import "github.com/vikash-parashar/bookings/internal/forms"

// Customized TemplateData To Send Data Into Templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
