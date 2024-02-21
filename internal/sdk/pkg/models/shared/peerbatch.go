// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/NetBird/terraform-provider-netbird/v4/internal/sdk/pkg/utils"
	"time"
)

type PeerBatch struct {
	// Number of accessible peers
	AccessiblePeersCount int64 `json:"accessible_peers_count"`
	// (Cloud only) Indicates whether peer needs approval
	ApprovalRequired *bool `json:"approval_required,omitempty"`
	// Commonly used English name of the city
	CityName string `json:"city_name"`
	// Peer to Management connection status
	Connected bool `json:"connected"`
	// Peer's public connection IP address
	ConnectionIP string `json:"connection_ip"`
	// 2-letter ISO 3166-1 alpha-2 code that represents the country
	CountryCode string `json:"country_code"`
	// Peer's DNS label is the parsed peer name for domain resolution. It is used to form an FQDN by appending the account's domain to the peer label. e.g. peer-dns-label.netbird.cloud
	DNSLabel string `json:"dns_label"`
	// Unique identifier from the GeoNames database for a specific geographical location.
	GeonameID int64 `json:"geoname_id"`
	// Groups that the peer belongs to
	Groups []GroupMinimum `json:"groups"`
	// Hostname of the machine
	Hostname string `json:"hostname"`
	// Peer ID
	ID string `json:"id"`
	// Peer's IP address
	IP string `json:"ip"`
	// Peer's operating system kernel version
	KernelVersion string `json:"kernel_version"`
	// Last time this peer performed log in (authentication). E.g., user authenticated.
	LastLogin time.Time `json:"last_login"`
	// Last time peer connected to Netbird's management service
	LastSeen time.Time `json:"last_seen"`
	// Indicates whether peer login expiration has been enabled or not
	LoginExpirationEnabled bool `json:"login_expiration_enabled"`
	// Indicates whether peer's login expired or not
	LoginExpired bool `json:"login_expired"`
	// Peer's hostname
	Name string `json:"name"`
	// Peer's operating system and version
	Os string `json:"os"`
	// Indicates whether SSH server is enabled on this peer
	SSHEnabled bool `json:"ssh_enabled"`
	// Peer's desktop UI version
	UIVersion string `json:"ui_version"`
	// User ID of the user that enrolled this peer
	UserID string `json:"user_id"`
	// Peer's daemon or cli version
	Version string `json:"version"`
}

func (p PeerBatch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PeerBatch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PeerBatch) GetAccessiblePeersCount() int64 {
	if o == nil {
		return 0
	}
	return o.AccessiblePeersCount
}

func (o *PeerBatch) GetApprovalRequired() *bool {
	if o == nil {
		return nil
	}
	return o.ApprovalRequired
}

func (o *PeerBatch) GetCityName() string {
	if o == nil {
		return ""
	}
	return o.CityName
}

func (o *PeerBatch) GetConnected() bool {
	if o == nil {
		return false
	}
	return o.Connected
}

func (o *PeerBatch) GetConnectionIP() string {
	if o == nil {
		return ""
	}
	return o.ConnectionIP
}

func (o *PeerBatch) GetCountryCode() string {
	if o == nil {
		return ""
	}
	return o.CountryCode
}

func (o *PeerBatch) GetDNSLabel() string {
	if o == nil {
		return ""
	}
	return o.DNSLabel
}

func (o *PeerBatch) GetGeonameID() int64 {
	if o == nil {
		return 0
	}
	return o.GeonameID
}

func (o *PeerBatch) GetGroups() []GroupMinimum {
	if o == nil {
		return []GroupMinimum{}
	}
	return o.Groups
}

func (o *PeerBatch) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *PeerBatch) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PeerBatch) GetIP() string {
	if o == nil {
		return ""
	}
	return o.IP
}

func (o *PeerBatch) GetKernelVersion() string {
	if o == nil {
		return ""
	}
	return o.KernelVersion
}

func (o *PeerBatch) GetLastLogin() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastLogin
}

func (o *PeerBatch) GetLastSeen() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastSeen
}

func (o *PeerBatch) GetLoginExpirationEnabled() bool {
	if o == nil {
		return false
	}
	return o.LoginExpirationEnabled
}

func (o *PeerBatch) GetLoginExpired() bool {
	if o == nil {
		return false
	}
	return o.LoginExpired
}

func (o *PeerBatch) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PeerBatch) GetOs() string {
	if o == nil {
		return ""
	}
	return o.Os
}

func (o *PeerBatch) GetSSHEnabled() bool {
	if o == nil {
		return false
	}
	return o.SSHEnabled
}

func (o *PeerBatch) GetUIVersion() string {
	if o == nil {
		return ""
	}
	return o.UIVersion
}

func (o *PeerBatch) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *PeerBatch) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}
