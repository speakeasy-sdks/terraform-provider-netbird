// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type PostAPIRoutesSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type PostAPIRoutesResponse struct {
	ContentType string
	// A Route Object
	Route       *shared.Route
	StatusCode  int
	RawResponse *http.Response
}