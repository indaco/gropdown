// Package gropdown - Helpers for using dropdown with `template/html`
package gropdown

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"html/template"
)

// HTMLGenerator provides functions for generating HTML code for dropdown.
type HTMLGenerator struct{}

// NewHTMLGenerator creates a new instance of HTMLGenerator.
func NewHTMLGenerator() *HTMLGenerator {
	return &HTMLGenerator{}
}

// GropdownCSSToGoHTML generates HTML code for the dropdown CSS and returns it as a template.HTML.
func (g *HTMLGenerator) GropdownCSSToGoHTML() (template.HTML, error) {
	html, err := templ.ToGoHTML(context.Background(), GropdownCSS())
	if err != nil {
		return "", fmt.Errorf("failed to generate dropdown CSS: %v", err)
	}
	return html, nil
}

// GropdownJSToGoHTML generates HTML code for the dropdown JS and returns it as a template.HTML.
func (g *HTMLGenerator) GropdownJSToGoHTML(dropdown *Dropdown) (template.HTML, error) {
	html, err := templ.ToGoHTML(context.Background(), GropdownJS(dropdown))
	if err != nil {
		return "", fmt.Errorf("failed to generate dropdown JS: %v", err)
	}
	return html, nil
}

// Render generates HTML code for displaying the dropdown and returns it as a template.HTML.
func (g *HTMLGenerator) Render(dd *DropdownBuilder) (template.HTML, error) {
	// Generate HTML code for displaying the toast.
	html, err := templ.ToGoHTML(context.Background(), dd.Render())
	if err != nil {
		return "", fmt.Errorf("failed to generate dropdown HTML: %v", err)
	}
	return html, nil
}
