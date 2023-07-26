// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Rule - A Rule Object
type Rule struct {
	// Rule friendly description
	Description string `json:"description"`
	// Rule destination groups
	Destinations []GroupMinimum `json:"destinations"`
	// Rules status
	Disabled bool `json:"disabled"`
	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`
	// Rule ID
	ID string `json:"id"`
	// Rule name identifier
	Name string `json:"name"`
	// Rule source groups
	Sources []GroupMinimum `json:"sources"`
}
