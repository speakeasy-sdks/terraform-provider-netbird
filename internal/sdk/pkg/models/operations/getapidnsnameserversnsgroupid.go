// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type GetAPIDNSNameserversNsgroupIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type GetAPIDNSNameserversNsgroupIDRequest struct {
	// The unique identifier of a Nameserver Group
	NsgroupID string `pathParam:"style=simple,explode=false,name=nsgroupId"`
}

type GetAPIDNSNameserversNsgroupIDResponse struct {
	ContentType string
	// A Nameserver Group object
	NameserverGroup *shared.NameserverGroup
	StatusCode      int
	RawResponse     *http.Response
}