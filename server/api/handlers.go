package api

import (
	"fmt"
	"net/http"
)

// Index is a handler for: /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
