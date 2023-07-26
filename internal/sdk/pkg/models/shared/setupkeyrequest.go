// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SetupKeyRequest - New Setup Key request
type SetupKeyRequest struct {
	// Setup key groups to auto-assign to peers registered with this key
	AutoGroups []string `json:"auto_groups"`
	// Expiration time in seconds
	ExpiresIn int64 `json:"expires_in"`
	// Setup Key name
	Name string `json:"name"`
	// Setup key revocation status
	Revoked bool `json:"revoked"`
	// Setup key type, one-off for single time usage and reusable
	Type string `json:"type"`
	// A number of times this key can be used. The value of 0 indicates the unlimited usage.
	UsageLimit int64 `json:"usage_limit"`
}
