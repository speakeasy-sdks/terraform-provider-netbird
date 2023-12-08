// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/NetBird/terraform-provider-netbird/v2/internal/sdk/pkg/utils"
	"time"
)

// Status - User's status
type Status string

const (
	StatusActive  Status = "active"
	StatusInvited Status = "invited"
	StatusBlocked Status = "blocked"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
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
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type User struct {
	// Group IDs to auto-assign to peers registered by this user
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
	// How user was issued by API or Integration
	Issued *string `json:"issued,omitempty"`
	// Last time this user performed a login to the dashboard
	LastLogin *time.Time `json:"last_login,omitempty"`
	// User's name from idp provider
	Name string `json:"name"`
	// User's NetBird account role
	Role string `json:"role"`
	// User's status
	Status Status `json:"status"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *User) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *User) GetAutoGroups() []string {
	if o == nil {
		return []string{}
	}
	return o.AutoGroups
}

func (o *User) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *User) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *User) GetIsBlocked() bool {
	if o == nil {
		return false
	}
	return o.IsBlocked
}

func (o *User) GetIsCurrent() *bool {
	if o == nil {
		return nil
	}
	return o.IsCurrent
}

func (o *User) GetIsServiceUser() *bool {
	if o == nil {
		return nil
	}
	return o.IsServiceUser
}

func (o *User) GetIssued() *string {
	if o == nil {
		return nil
	}
	return o.Issued
}

func (o *User) GetLastLogin() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastLogin
}

func (o *User) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *User) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *User) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}
