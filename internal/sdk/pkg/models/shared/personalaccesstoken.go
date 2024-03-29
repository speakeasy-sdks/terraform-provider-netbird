// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/NetBird/terraform-provider-netbird/v4/internal/sdk/pkg/utils"
	"time"
)

type PersonalAccessToken struct {
	// Date the token was created
	CreatedAt time.Time `json:"created_at"`
	// User ID of the user who created the token
	CreatedBy string `json:"created_by"`
	// Date the token expires
	ExpirationDate time.Time `json:"expiration_date"`
	// ID of a token
	ID string `json:"id"`
	// Date the token was last used
	LastUsed *time.Time `json:"last_used,omitempty"`
	// Name of the token
	Name string `json:"name"`
}

func (p PersonalAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonalAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonalAccessToken) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PersonalAccessToken) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *PersonalAccessToken) GetExpirationDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ExpirationDate
}

func (o *PersonalAccessToken) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PersonalAccessToken) GetLastUsed() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUsed
}

func (o *PersonalAccessToken) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
