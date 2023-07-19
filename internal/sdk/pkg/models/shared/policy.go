// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Policy - A Policy Object
type Policy struct {
	// Policy friendly description
	Description string `json:"description"`
	// Policy status
	Enabled bool `json:"enabled"`
	// Policy ID
	ID *string `json:"id,omitempty"`
	// Policy name identifier
	Name string `json:"name"`
	// Policy Rego query
	Query string `json:"query"`
	// Policy rule object for policy UI editor
	Rules []PolicyRule `json:"rules"`
}