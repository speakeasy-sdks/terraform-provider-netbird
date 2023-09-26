// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type GetAPIPoliciesPolicyIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type GetAPIPoliciesPolicyIDRequest struct {
	// The unique identifier of a policy
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
}

type GetAPIPoliciesPolicyIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A Policy object
	Policy *shared.Policy
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
