// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountExtraSettings struct {
	// (Cloud only) Enables or disables peer approval globally. If enabled, all peers added will be in pending state until approved by an admin.
	PeerApprovalEnabled *bool `json:"peer_approval_enabled,omitempty"`
}

func (o *AccountExtraSettings) GetPeerApprovalEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PeerApprovalEnabled
}
