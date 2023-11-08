// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/NetBird/terraform-provider-netbird/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutAPIRulesRuleIDSecurity struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	TokenAuth  *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *PutAPIRulesRuleIDSecurity) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}

func (o *PutAPIRulesRuleIDSecurity) GetTokenAuth() *string {
	if o == nil {
		return nil
	}
	return o.TokenAuth
}

type PutAPIRulesRuleIDRequest struct {
	// Update Rule request
	RuleRequest *shared.RuleRequest `request:"mediaType=application/json"`
	// The unique identifier of a rule
	RuleID string `pathParam:"style=simple,explode=false,name=ruleId"`
}

func (o *PutAPIRulesRuleIDRequest) GetRuleRequest() *shared.RuleRequest {
	if o == nil {
		return nil
	}
	return o.RuleRequest
}

func (o *PutAPIRulesRuleIDRequest) GetRuleID() string {
	if o == nil {
		return ""
	}
	return o.RuleID
}

type PutAPIRulesRuleIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A Rule object
	Rule *shared.Rule
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutAPIRulesRuleIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAPIRulesRuleIDResponse) GetRule() *shared.Rule {
	if o == nil {
		return nil
	}
	return o.Rule
}

func (o *PutAPIRulesRuleIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAPIRulesRuleIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
