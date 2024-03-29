// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/NetBird/terraform-provider-netbird/v4/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetAPIPostureChecksSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *GetAPIPostureChecksSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *GetAPIPostureChecksSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type GetAPIPostureChecksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A JSON Array of posture checks
	Classes []shared.PostureCheck
}

func (o *GetAPIPostureChecksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIPostureChecksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIPostureChecksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPIPostureChecksResponse) GetClasses() []shared.PostureCheck {
	if o == nil {
		return nil
	}
	return o.Classes
}
