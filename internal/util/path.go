// internal/util/path.go
package util

import (
	"net/http"
	"path/filepath"
	"strings"
)

func CleanPath(path string) string {
	return filepath.Clean(strings.TrimSpace(path))
}

func IsHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
