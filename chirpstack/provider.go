package chirpstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc"
)

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	ConnectionData ConnectionData

	diagnostics diag.Diagnostics
	ctx         context.Context
	conn        *grpc.ClientConn
}

func (p *provider) Diagnostics() diag.Diagnostics {
	diags := p.diagnostics
	p.diagnostics = []diag.Diagnostic{}
	return diags
}

// GetSchema
func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"server": {
				Type:     types.StringType,
				Optional: true,
			},
			"key": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
			"email": {
				Type:     types.StringType,
				Optional: true,
			},
			"password": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
		},
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config ConnectionData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	p.ConnectionData = config
}

// GetResources - Defines provider resources
func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"chirpstack_network-server":    resourceNetworkServerType{},
		"chirpstack_api-key":           resourceAPIKeyType{},
		"chirpstack_user":              resourceUserType{},
		"chirpstack_organization":      resourceOrganizationType{},
		"chirpstack_organization-user": resourceOrganizationUserType{},
		"chirpstack_service-profile":   resourceServiceProfileType{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
