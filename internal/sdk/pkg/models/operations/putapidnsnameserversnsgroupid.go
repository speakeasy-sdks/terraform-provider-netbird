// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/NetBird/terraform-provider-netbird/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutAPIDNSNameserversNsgroupIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *PutAPIDNSNameserversNsgroupIDSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *PutAPIDNSNameserversNsgroupIDSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type PutAPIDNSNameserversNsgroupIDRequest struct {
	// Update Nameserver Group request
	NameserverGroupRequest *shared.NameserverGroupRequest `request:"mediaType=application/json"`
	// The unique identifier of a Nameserver Group
	NsgroupID string `pathParam:"style=simple,explode=false,name=nsgroupId"`
}

func (o *PutAPIDNSNameserversNsgroupIDRequest) GetNameserverGroupRequest() *shared.NameserverGroupRequest {
	if o == nil {
		return nil
	}
	return o.NameserverGroupRequest
}

func (o *PutAPIDNSNameserversNsgroupIDRequest) GetNsgroupID() string {
	if o == nil {
		return ""
	}
	return o.NsgroupID
}

type PutAPIDNSNameserversNsgroupIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A Nameserver Group object
	NameserverGroup *shared.NameserverGroup
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutAPIDNSNameserversNsgroupIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAPIDNSNameserversNsgroupIDResponse) GetNameserverGroup() *shared.NameserverGroup {
	if o == nil {
		return nil
	}
	return o.NameserverGroup
}

func (o *PutAPIDNSNameserversNsgroupIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAPIDNSNameserversNsgroupIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
