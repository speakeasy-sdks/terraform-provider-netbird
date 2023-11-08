// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/NetBird/terraform-provider-netbird/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostAPIUsersUserIDTokensSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *PostAPIUsersUserIDTokensSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *PostAPIUsersUserIDTokensSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type PostAPIUsersUserIDTokensRequest struct {
	// PersonalAccessToken create parameters
	PersonalAccessTokenRequest *shared.PersonalAccessTokenRequest `request:"mediaType=application/json"`
	// The unique identifier of a user
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *PostAPIUsersUserIDTokensRequest) GetPersonalAccessTokenRequest() *shared.PersonalAccessTokenRequest {
	if o == nil {
		return nil
	}
	return o.PersonalAccessTokenRequest
}

func (o *PostAPIUsersUserIDTokensRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type PostAPIUsersUserIDTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The token in plain text
	PersonalAccessTokenGenerated *shared.PersonalAccessTokenGenerated
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPIUsersUserIDTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIUsersUserIDTokensResponse) GetPersonalAccessTokenGenerated() *shared.PersonalAccessTokenGenerated {
	if o == nil {
		return nil
	}
	return o.PersonalAccessTokenGenerated
}

func (o *PostAPIUsersUserIDTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIUsersUserIDTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
