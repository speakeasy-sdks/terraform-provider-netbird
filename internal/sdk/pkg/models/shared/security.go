// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}
