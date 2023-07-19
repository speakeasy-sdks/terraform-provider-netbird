// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type PutAPIPoliciesPolicyIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type PutAPIPoliciesPolicyIDRequest struct {
	// Update Policy request
	PolicyUpdate *shared.PolicyUpdate `request:"mediaType=application/json"`
	// The unique identifier of a policy
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
}

type PutAPIPoliciesPolicyIDResponse struct {
	ContentType string
	// A Policy object
	Policy      *shared.Policy
	StatusCode  int
	RawResponse *http.Response
}