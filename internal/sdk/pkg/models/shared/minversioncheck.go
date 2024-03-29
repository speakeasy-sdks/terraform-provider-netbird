// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MinVersionCheck - Posture check for the version of operating system
type MinVersionCheck struct {
	// Minimum acceptable version
	MinVersion string `json:"min_version"`
}

func (o *MinVersionCheck) GetMinVersion() string {
	if o == nil {
		return ""
	}
	return o.MinVersion
}
