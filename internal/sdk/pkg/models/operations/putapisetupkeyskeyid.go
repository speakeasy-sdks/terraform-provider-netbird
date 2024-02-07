// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/NetBird/terraform-provider-netbird/v4/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutAPISetupKeysKeyIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *PutAPISetupKeysKeyIDSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *PutAPISetupKeysKeyIDSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type PutAPISetupKeysKeyIDRequest struct {
	// update to Setup Key
	SetupKeyRequest *shared.SetupKeyRequest `request:"mediaType=application/json"`
	// The unique identifier of a setup key
	KeyID string `pathParam:"style=simple,explode=false,name=keyId"`
}

func (o *PutAPISetupKeysKeyIDRequest) GetSetupKeyRequest() *shared.SetupKeyRequest {
	if o == nil {
		return nil
	}
	return o.SetupKeyRequest
}

func (o *PutAPISetupKeysKeyIDRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

type PutAPISetupKeysKeyIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A Setup Key object
	SetupKey *shared.SetupKey
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutAPISetupKeysKeyIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAPISetupKeysKeyIDResponse) GetSetupKey() *shared.SetupKey {
	if o == nil {
		return nil
	}
	return o.SetupKey
}

func (o *PutAPISetupKeysKeyIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAPISetupKeysKeyIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
