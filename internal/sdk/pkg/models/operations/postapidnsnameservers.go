// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type PostAPIDNSNameserversSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type PostAPIDNSNameserversResponse struct {
	ContentType string
	// A Nameserver Groups Object
	NameserverGroup *shared.NameserverGroup
	StatusCode      int
	RawResponse     *http.Response
}