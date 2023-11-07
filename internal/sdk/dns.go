// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"netbird/v2/internal/sdk/pkg/models/operations"
	"netbird/v2/internal/sdk/pkg/models/sdkerrors"
	"netbird/v2/internal/sdk/pkg/models/shared"
	"netbird/v2/internal/sdk/pkg/utils"
	"strings"
)

// DNS - Interact with and view information about DNS configuration.
type DNS struct {
	sdkConfiguration sdkConfiguration
}

func newDNS(sdkConfig sdkConfiguration) *DNS {
	return &DNS{
		sdkConfiguration: sdkConfig,
	}
}

// DeleteAPIDNSNameserversNsgroupID - Delete a Nameserver Group
// Delete a Nameserver Group
func (s *DNS) DeleteAPIDNSNameserversNsgroupID(ctx context.Context, request operations.DeleteAPIDNSNameserversNsgroupIDRequest, security operations.DeleteAPIDNSNameserversNsgroupIDSecurity) (*operations.DeleteAPIDNSNameserversNsgroupIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/dns/nameservers/{nsgroupId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteAPIDNSNameserversNsgroupIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		fallthrough
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// GetAPIDNSNameservers - List all Nameserver Groups
// Returns a list of all Nameserver Groups
func (s *DNS) GetAPIDNSNameservers(ctx context.Context, security operations.GetAPIDNSNameserversSecurity) (*operations.GetAPIDNSNameserversResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/dns/nameservers"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIDNSNameserversResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.NameserverGroup
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Classes = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// GetAPIDNSNameserversNsgroupID - Retrieve a Nameserver Group
// Get information about a Nameserver Groups
func (s *DNS) GetAPIDNSNameserversNsgroupID(ctx context.Context, request operations.GetAPIDNSNameserversNsgroupIDRequest, security operations.GetAPIDNSNameserversNsgroupIDSecurity) (*operations.GetAPIDNSNameserversNsgroupIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/dns/nameservers/{nsgroupId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIDNSNameserversNsgroupIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.NameserverGroup
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.NameserverGroup = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// GetAPIDNSSettings - Retrieve DNS settings
// Returns a DNS settings object
func (s *DNS) GetAPIDNSSettings(ctx context.Context, security operations.GetAPIDNSSettingsSecurity) (*operations.GetAPIDNSSettingsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/dns/settings"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIDNSSettingsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.DNSSettings
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Classes = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// PostAPIDNSNameservers - Create a Nameserver Group
// Creates a Nameserver Group
func (s *DNS) PostAPIDNSNameservers(ctx context.Context, request *shared.NameserverGroupRequest, security operations.PostAPIDNSNameserversSecurity) (*operations.PostAPIDNSNameserversResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/dns/nameservers"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIDNSNameserversResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.NameserverGroup
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.NameserverGroup = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// PutAPIDNSNameserversNsgroupID - Update a Nameserver Group
// Update/Replace a Nameserver Group
func (s *DNS) PutAPIDNSNameserversNsgroupID(ctx context.Context, request operations.PutAPIDNSNameserversNsgroupIDRequest, security operations.PutAPIDNSNameserversNsgroupIDSecurity) (*operations.PutAPIDNSNameserversNsgroupIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/dns/nameservers/{nsgroupId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "NameserverGroupRequest", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutAPIDNSNameserversNsgroupIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.NameserverGroup
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.NameserverGroup = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// PutAPIDNSSettings - Update DNS Settings
// Updates a DNS settings object
func (s *DNS) PutAPIDNSSettings(ctx context.Context, request *shared.DNSSettings, security operations.PutAPIDNSSettingsSecurity) (*operations.PutAPIDNSSettingsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/dns/settings"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutAPIDNSSettingsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.DNSSettings
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.DNSSettings = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}
