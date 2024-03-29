// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/NetBird/terraform-provider-netbird/v4/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutAPIPeersPeerIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *PutAPIPeersPeerIDSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *PutAPIPeersPeerIDSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type PutAPIPeersPeerIDRequest struct {
	// update a peer
	PeerRequest *shared.PeerRequest `request:"mediaType=application/json"`
	// The unique identifier of a peer
	PeerID string `pathParam:"style=simple,explode=false,name=peerId"`
}

func (o *PutAPIPeersPeerIDRequest) GetPeerRequest() *shared.PeerRequest {
	if o == nil {
		return nil
	}
	return o.PeerRequest
}

func (o *PutAPIPeersPeerIDRequest) GetPeerID() string {
	if o == nil {
		return ""
	}
	return o.PeerID
}

type PutAPIPeersPeerIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A Peer object
	Peer *shared.Peer
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutAPIPeersPeerIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAPIPeersPeerIDResponse) GetPeer() *shared.Peer {
	if o == nil {
		return nil
	}
	return o.Peer
}

func (o *PutAPIPeersPeerIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAPIPeersPeerIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
