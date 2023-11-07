// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"netbird/v2/internal/sdk/pkg/utils"
	"time"
)

type SetupKey struct {
	// List of group IDs to auto-assign to peers registered with this key
	AutoGroups []string `json:"auto_groups"`
	// Indicate that the peer will be ephemeral or not
	Ephemeral bool `json:"ephemeral"`
	// Setup Key expiration date
	Expires time.Time `json:"expires"`
	// Setup Key ID
	ID string `json:"id"`
	// Setup Key value
	Key string `json:"key"`
	// Setup key last usage date
	LastUsed time.Time `json:"last_used"`
	// Setup key name identifier
	Name string `json:"name"`
	// Setup key revocation status
	Revoked bool `json:"revoked"`
	// Setup key status, "valid", "overused","expired" or "revoked"
	State string `json:"state"`
	// Setup key type, one-off for single time usage and reusable
	Type string `json:"type"`
	// Setup key last update date
	UpdatedAt time.Time `json:"updated_at"`
	// A number of times this key can be used. The value of 0 indicates the unlimited usage.
	UsageLimit int64 `json:"usage_limit"`
	// Usage count of setup key
	UsedTimes int64 `json:"used_times"`
	// Setup key validity status
	Valid bool `json:"valid"`
}

func (s SetupKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SetupKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SetupKey) GetAutoGroups() []string {
	if o == nil {
		return []string{}
	}
	return o.AutoGroups
}

func (o *SetupKey) GetEphemeral() bool {
	if o == nil {
		return false
	}
	return o.Ephemeral
}

func (o *SetupKey) GetExpires() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Expires
}

func (o *SetupKey) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SetupKey) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *SetupKey) GetLastUsed() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastUsed
}

func (o *SetupKey) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SetupKey) GetRevoked() bool {
	if o == nil {
		return false
	}
	return o.Revoked
}

func (o *SetupKey) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *SetupKey) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *SetupKey) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *SetupKey) GetUsageLimit() int64 {
	if o == nil {
		return 0
	}
	return o.UsageLimit
}

func (o *SetupKey) GetUsedTimes() int64 {
	if o == nil {
		return 0
	}
	return o.UsedTimes
}

func (o *SetupKey) GetValid() bool {
	if o == nil {
		return false
	}
	return o.Valid
}
