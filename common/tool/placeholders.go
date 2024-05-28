package tool

import "strings"

// Construct n anonymous parameter placeholders use for sql
func InPlaceholders(n int) string {
	var b strings.Builder
	for i := 0; i < n - 1; i ++ {
		b.WriteString("?,")
	}
	if n > 0 {
		b.WriteString("?")
	}
	return b.String()
}