package paging

import "github.com/nkpetko/goblox/rbx/model/shared"

// APIPageResponse represents a page of results.
// This doesn't actually contain the results, but instead it has the other paging parameters such as cursors.
// Inherited structs should implement the results field.
type APIPageResponse struct {
	// The raw response from the Roblox API.
	// ignore this field with json, yaml, etc.
	shared.HTTPSharedModel `json:"-" yaml:"-"`

	// The cursor for the previous page
	PreviousPageCursor string `json:"previousPageCursor"`

	// The cursor for the next page
	NextPageCursor string `json:"nextPageCursor"`
}
