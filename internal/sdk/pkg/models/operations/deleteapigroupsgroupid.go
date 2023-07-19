// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAPIGroupsGroupIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

type DeleteAPIGroupsGroupIDRequest struct {
	// The unique identifier of a group
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
}

type DeleteAPIGroupsGroupIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}