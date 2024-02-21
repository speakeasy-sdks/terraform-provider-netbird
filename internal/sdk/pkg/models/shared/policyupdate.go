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
	// Policy rule object for policy UI editor
	Rules []PolicyRuleUpdate `json:"rules"`
	// Posture checks ID's applied to policy source groups
	SourcePostureChecks []string `json:"source_posture_checks,omitempty"`
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

func (o *PolicyUpdate) GetRules() []PolicyRuleUpdate {
	if o == nil {
		return []PolicyRuleUpdate{}
	}
	return o.Rules
}

func (o *PolicyUpdate) GetSourcePostureChecks() []string {
	if o == nil {
		return nil
	}
	return o.SourcePostureChecks
}
