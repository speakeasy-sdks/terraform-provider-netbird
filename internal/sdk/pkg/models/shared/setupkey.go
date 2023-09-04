// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// SetupKey - A Setup Keys Object
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
