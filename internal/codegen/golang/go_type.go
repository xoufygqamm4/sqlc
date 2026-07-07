package golang

import (
	"github.com/sqlc-dev/sqlc/internal/plugin"
)

func goType(col *plugin.Column, settings Settings) string {
	// ... existing code ...
	if col.IsArray {
		colCopy := *col
		colCopy.IsArray = false
		colCopy.NotNull = true
		elemType := goType(&colCopy, settings)
		if !col.NotNull {
			return "*[]" + elemType
		}
		return "[]" + elemType
	}
	// ... rest of existing code ...
}