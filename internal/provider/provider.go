// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"netbird/v2/internal/sdk"
	"netbird/v2/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &NetbirdProvider{}

type NetbirdProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NetbirdProviderModel describes the provider data model.
type NetbirdProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
	TokenAuth  types.String `tfsdk:"token_auth"`
}

func (p *NetbirdProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netbird"
	resp.Version = p.version
}

func (p *NetbirdProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `NetBird REST API: API to manipulate groups, rules, policies and retrieve information about peers and users`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.netbird.io)",
				Optional:            true,
				Required:            false,
			},
			"bearer_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"token_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *NetbirdProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data NetbirdProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.netbird.io"
	}

	bearerAuth := new(string)
	if !data.BearerAuth.IsUnknown() && !data.BearerAuth.IsNull() {
		*bearerAuth = data.BearerAuth.ValueString()
	} else {
		bearerAuth = nil
	}
	tokenAuth := new(string)
	if !data.TokenAuth.IsUnknown() && !data.TokenAuth.IsNull() {
		*tokenAuth = data.TokenAuth.ValueString()
	} else {
		tokenAuth = nil
	}
	security := shared.Security{
		BearerAuth: bearerAuth,
		TokenAuth:  tokenAuth,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *NetbirdProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *NetbirdProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NetbirdProvider{
			version: version,
		}
	}
}
