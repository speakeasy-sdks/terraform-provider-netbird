// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/NetBird/terraform-provider-netbird/v4/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetAPIUsersUserIDTokensTokenIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *GetAPIUsersUserIDTokensTokenIDSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *GetAPIUsersUserIDTokensTokenIDSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type GetAPIUsersUserIDTokensTokenIDRequest struct {
	// The unique identifier of a token
	TokenID string `pathParam:"style=simple,explode=false,name=tokenId"`
	// The unique identifier of a user
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *GetAPIUsersUserIDTokensTokenIDRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

func (o *GetAPIUsersUserIDTokensTokenIDRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type GetAPIUsersUserIDTokensTokenIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A PersonalAccessTokens Object
	PersonalAccessToken *shared.PersonalAccessToken
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAPIUsersUserIDTokensTokenIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIUsersUserIDTokensTokenIDResponse) GetPersonalAccessToken() *shared.PersonalAccessToken {
	if o == nil {
		return nil
	}
	return o.PersonalAccessToken
}

func (o *GetAPIUsersUserIDTokensTokenIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIUsersUserIDTokensTokenIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
