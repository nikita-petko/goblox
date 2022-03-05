package shared

// APIEmptyResponseModel is used to represent an empty model.
// It is used for documetation, not anything else.
//
// This represents the following model:
// Roblox.Web.WebAPI.APIEmptyResponseModel
type APIEmptyResponseModel struct {
	// The raw response from the Roblox API.
	// ignore this field with json, yaml, etc.
	HTTPSharedModel `json:"-" yaml:"-"`
}
