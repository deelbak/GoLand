package models

import "github.com/deelbak/go-project-kbtu/internal/forms"

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	Form            *forms.Form
	IsAuthenticated int
}
