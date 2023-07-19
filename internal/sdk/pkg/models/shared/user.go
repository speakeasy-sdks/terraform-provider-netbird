// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UserStatus - User's status
type UserStatus string

const (
	UserStatusActive  UserStatus = "active"
	UserStatusInvited UserStatus = "invited"
	UserStatusBlocked UserStatus = "blocked"
)

func (e UserStatus) ToPointer() *UserStatus {
	return &e
}

func (e *UserStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "invited":
		fallthrough
	case "blocked":
		*e = UserStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserStatus: %v", v)
	}
}

// User - A User object
type User struct {
	// Groups to auto-assign to peers registered by this user
	AutoGroups []string `json:"auto_groups"`
	// User's email address
	Email string `json:"email"`
	// User ID
	ID string `json:"id"`
	// Is true if this user is blocked. Blocked users can't use the system
	IsBlocked bool `json:"is_blocked"`
	// Is true if authenticated user is the same as this user
	IsCurrent *bool `json:"is_current,omitempty"`
	// Is true if this user is a service user
	IsServiceUser *bool `json:"is_service_user,omitempty"`
	// User's name from idp provider
	Name string `json:"name"`
	// User's NetBird account role
	Role string `json:"role"`
	// User's status
	Status UserStatus `json:"status"`
}