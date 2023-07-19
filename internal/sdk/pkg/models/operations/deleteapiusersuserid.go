// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAPIUsersUserIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type DeleteAPIUsersUserIDRequest struct {
	// The unique identifier of a user
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type DeleteAPIUsersUserIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}