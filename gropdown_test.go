package gropdown

import (
	"testing"

	"github.com/a-h/templ"
)

func TestDropdownBuilder(t *testing.T) {
	tests := []struct {
		name                  string
		builder               *DropdownBuilder
		expectedOpenState     bool
		expectedAnimation     bool
		expectedPosition      Position
		expectedButtonLabel   string
		expectedButtonIcon    string
		expectedButtonAttrs   map[string]string
		expectedItems         []DropdownItem
		expectedDividerExists bool
	}{
		// Test cases for options and config settings
		{
			name: "Dropdown with Open State Set to True",
			builder: NewDropdownBuilder().
				SetOpen(true).
				SetAnimation(false).
				SetPosition(Bottom).
				WithButton(DropdownButton{
					Label: "Dropdown",
					Icon:  "icon",
				}).
				WithItems([]DropdownItem{
					{Label: "Item 1"},
					{Label: "Item 2"},
					{Label: "Item 3"},
				}),
			expectedOpenState: true,
			expectedPosition:  Bottom,
		},
		{
			name: "Dropdown with Animation Disabled",
			builder: NewDropdownBuilder().
				SetOpen(false).
				SetAnimation(false).
				SetPosition(Bottom).
				WithButton(DropdownButton{
					Label: "Dropdown",
					Icon:  "icon",
				}).
				WithItems([]DropdownItem{
					{Label: "Item 1"},
					{Label: "Item 2"},
					{Label: "Item 3"},
				}),
			expectedAnimation: false,
			expectedPosition:  Bottom,
		},
		{
			name: "Dropdown with Custom Button Configuration",
			builder: NewDropdownBuilder().
				SetOpen(false).
				SetAnimation(false).
				SetPosition(Left).
				WithButton(DropdownButton{
					Label: "Custom Button",
					Icon:  "custom-icon",
					Attrs: templ.Attributes{
						"id":    "custom-button-id",
						"class": "custom-button-class",
					},
				}).
				WithItems([]DropdownItem{
					{Label: "Item 1"},
					{Label: "Item 2"},
					{Label: "Item 3"},
				}),
			expectedPosition: Left,
		},
		{
			name: "Dropdown with Custom Item Configuration",
			builder: NewDropdownBuilder().
				SetOpen(false).
				SetAnimation(false).
				SetPosition(Right).
				WithButton(DropdownButton{
					Label: "Dropdown",
					Icon:  "icon",
				}).
				WithItems([]DropdownItem{
					{Label: "Item 1", Icon: "item-icon-1", Href: "/item1", External: true},
					{Label: "Item 2", Icon: "item-icon-2", Href: "/item2"},
					{Label: "Item 3", Icon: "item-icon-3"},
				}),
			expectedPosition: Right,
		},
		{
			name: "Dropdown with Divider",
			builder: NewDropdownBuilder().
				SetOpen(false).
				SetAnimation(false).
				SetPosition(Right).
				WithButton(DropdownButton{
					Label: "Dropdown",
					Icon:  "icon",
				}).
				WithItems([]DropdownItem{
					{Label: "Item 1", Icon: "item-icon-1", Href: "/item1", External: true},
					{Label: "Item 2", Icon: "item-icon-2", Href: "/item2"},
					{Label: "Item 3", Icon: "item-icon-3"},
					{},
					{Label: "Item 4", Icon: "item-icon-4", Href: "/item4"},
				}),
			expectedPosition:      Right,
			expectedDividerExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dropdown := tt.builder.Dropdown()

			if tt.expectedOpenState != dropdown.Open {
				t.Errorf("unexpected open state: got %v, want %v", dropdown.Open, tt.expectedOpenState)
			}

			if tt.expectedAnimation != dropdown.Animation {
				t.Errorf("unexpected animation setting: got %v, want %v", dropdown.Animation, tt.expectedAnimation)
			}

			if tt.expectedPosition != dropdown.Position {
				t.Errorf("unexpected position: got %v, want %v", dropdown.Position, tt.expectedPosition)
			}
		})
	}
}
