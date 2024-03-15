package gropdown

import (
	"fmt"
	"strings"
)

func buttonId(label string) string {
	return fmt.Sprintf("dropdown-button-%s", strings.ToLower(label))
}

func menuId(label string) string {
	return fmt.Sprintf("dropdown-menu-%s", strings.ToLower(label))
}
