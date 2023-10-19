// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"netbird/internal/sdk/pkg/models/shared"
)

type GetAPIRulesRuleIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *GetAPIRulesRuleIDSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *GetAPIRulesRuleIDSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type GetAPIRulesRuleIDRequest struct {
	// The unique identifier of a rule
	RuleID string `pathParam:"style=simple,explode=false,name=ruleId"`
}

func (o *GetAPIRulesRuleIDRequest) GetRuleID() string {
	if o == nil {
		return ""
	}
	return o.RuleID
}

type GetAPIRulesRuleIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A Rule object
	Rule *shared.Rule
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAPIRulesRuleIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIRulesRuleIDResponse) GetRule() *shared.Rule {
	if o == nil {
		return nil
	}
	return o.Rule
}

func (o *GetAPIRulesRuleIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIRulesRuleIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
