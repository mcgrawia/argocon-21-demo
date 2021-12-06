package workflows

import (
	"strings"
)

// DependsOnMultiple generates dependency string for the given template names or statements
func DependsOnMultiple(templates ...string) string {
	return strings.Join(templates, " && ")
}
