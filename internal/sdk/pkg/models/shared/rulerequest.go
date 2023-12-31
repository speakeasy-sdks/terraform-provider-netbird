// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RuleRequest struct {
	// Rule friendly description
	Description string `json:"description"`
	// List of destination group IDs
	Destinations []string `json:"destinations,omitempty"`
	// Rules status
	Disabled bool `json:"disabled"`
	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`
	// Rule name identifier
	Name string `json:"name"`
	// List of source group IDs
	Sources []string `json:"sources,omitempty"`
}

func (o *RuleRequest) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *RuleRequest) GetDestinations() []string {
	if o == nil {
		return nil
	}
	return o.Destinations
}

func (o *RuleRequest) GetDisabled() bool {
	if o == nil {
		return false
	}
	return o.Disabled
}

func (o *RuleRequest) GetFlow() string {
	if o == nil {
		return ""
	}
	return o.Flow
}

func (o *RuleRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RuleRequest) GetSources() []string {
	if o == nil {
		return nil
	}
	return o.Sources
}
