// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type PostAPIUsersSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type PostAPIUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A User object
	User *shared.User
}
