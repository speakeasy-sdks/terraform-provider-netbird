// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type GetAPISetupKeysSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type GetAPISetupKeysResponse struct {
	ContentType string
	// A JSON Array of Setup keys
	SetupKeys   []shared.SetupKey
	StatusCode  int
	RawResponse *http.Response
}
