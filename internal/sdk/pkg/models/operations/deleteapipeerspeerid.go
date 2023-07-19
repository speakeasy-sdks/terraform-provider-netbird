// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAPIPeersPeerIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type DeleteAPIPeersPeerIDRequest struct {
	// The unique identifier of a peer
	PeerID string `pathParam:"style=simple,explode=false,name=peerId"`
}

type DeleteAPIPeersPeerIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}