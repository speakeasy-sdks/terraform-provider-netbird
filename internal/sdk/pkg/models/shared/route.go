// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Route - A Route Object
type Route struct {
	// Route description
	Description string `json:"description"`
	// Route status
	Enabled bool `json:"enabled"`
	// Route group tag groups
	Groups []string `json:"groups"`
	// Route Id
	ID string `json:"id"`
	// Indicate if peer should masquerade traffic to this route's prefix
	Masquerade bool `json:"masquerade"`
	// Route metric number. Lowest number has higher priority
	Metric int64 `json:"metric"`
	// Network range in CIDR format
	Network string `json:"network"`
	// Route network identifier, to group HA routes
	NetworkID string `json:"network_id"`
	// Network type indicating if it is IPv4 or IPv6
	NetworkType string `json:"network_type"`
	// Peer Identifier associated with route
	Peer string `json:"peer"`
}