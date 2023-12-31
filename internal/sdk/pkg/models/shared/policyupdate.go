// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PolicyUpdate struct {
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
	Rules []PolicyRuleUpdate `json:"rules"`
}

func (o *PolicyUpdate) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *PolicyUpdate) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *PolicyUpdate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PolicyUpdate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PolicyUpdate) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *PolicyUpdate) GetRules() []PolicyRuleUpdate {
	if o == nil {
		return []PolicyRuleUpdate{}
	}
	return o.Rules
}
