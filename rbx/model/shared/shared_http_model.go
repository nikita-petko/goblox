package shared

import "net/http"

// HTTPSharedModel encapsulates the raw response from the Roblox API. Use this only if you want headers etc.
// This is a low level package, and should not be used by most users.
type HTTPSharedModel struct {
	// The raw response from the Roblox API.
	// ignore this field with json, yaml, etc.
	Response *http.Response `json:"-" yaml:"-"`
}
