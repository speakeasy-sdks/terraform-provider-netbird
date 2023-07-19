// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserRequest - User update
type UserRequest struct {
	// Groups to auto-assign to peers registered by this user
	AutoGroups []string `json:"auto_groups"`
	// If set to true then user is blocked and can't use the system
	IsBlocked bool `json:"is_blocked"`
	// User's NetBird account role
	Role string `json:"role"`
}