// internal/util/path.go
package util

import (
	"net/http"
	"path/filepath"
	"strings"
)

// CleanPath sanitizes and normalizes URL paths
func CleanPath(path string) string {
	return filepath.Clean(strings.TrimSpace(path))
}

// IsHtmxRequest checks if request is from HTMX
func IsHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
