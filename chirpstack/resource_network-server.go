package chirpstack

import (
	"context"
	"strconv"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceNetworkServerType struct{}

// NetworkServer Resource schema
func (r resourceNetworkServerType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.Int64Type,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"server": {
				Type:     types.StringType,
				Required: true,
			},
			"ca_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"tls_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"routing_profil_ca_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"routing_profil_tls_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"routing_profil_tls_key": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"gateway_discovery_enabled": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"gateway_discovery_interval": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"gateway_discovery_tx_frequency": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"gateway_discovery_dr": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceNetworkServerType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceNetworkServer{
		p: *(p.(*provider)),
	}, nil
}

type resourceNetworkServer struct {
	p provider
}

// Create a new resource
func (r resourceNetworkServer) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan NetworkServer
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	networkServer := api.NetworkServer{
		Name:                        plan.Name.Value,
		Server:                      plan.Server.Value,
		CaCert:                      plan.CACert.Value,
		TlsCert:                     plan.TLSCert.Value,
		RoutingProfileCaCert:        plan.RoutingProfileCACert.Value,
		RoutingProfileTlsCert:       plan.RoutingProfileTLSCert.Value,
		RoutingProfileTlsKey:        plan.RoutingProfileTLSKey.Value,
		GatewayDiscoveryEnabled:     plan.GatewayDiscoveryEnabled.Value,
		GatewayDiscoveryInterval:    uint32(plan.GatewayDiscoveryInterval.Value),
		GatewayDiscoveryTxFrequency: uint32(plan.GatewayDiscoveryTXFrequency.Value),
		GatewayDiscoveryDr:          uint32(plan.GatewayDiscoveryDR.Value),
	}
	request := api.CreateNetworkServerRequest{
		NetworkServer: &networkServer,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Create(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating network server",
			err.Error(),
		)
		return
	}

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), response.Id)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceNetworkServer) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state NetworkServer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetNetworkServerRequest{
		Id: state.ID.Value,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Get(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading network server",
			err.Error(),
		)
		return
	}

	state.Name = types.String{Value: response.NetworkServer.Name}
	state.Server = types.String{Value: response.NetworkServer.Server}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceNetworkServer) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan NetworkServer
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state NetworkServer
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = state.ID

	networkServer := api.NetworkServer{
		Id:                          plan.ID.Value,
		Name:                        plan.Name.Value,
		Server:                      plan.Server.Value,
		CaCert:                      plan.CACert.Value,
		TlsCert:                     plan.TLSCert.Value,
		RoutingProfileCaCert:        plan.RoutingProfileCACert.Value,
		RoutingProfileTlsCert:       plan.RoutingProfileTLSCert.Value,
		RoutingProfileTlsKey:        plan.RoutingProfileTLSKey.Value,
		GatewayDiscoveryEnabled:     plan.GatewayDiscoveryEnabled.Value,
		GatewayDiscoveryInterval:    uint32(plan.GatewayDiscoveryInterval.Value),
		GatewayDiscoveryTxFrequency: uint32(plan.GatewayDiscoveryTXFrequency.Value),
		GatewayDiscoveryDr:          uint32(plan.GatewayDiscoveryDR.Value),
	}
	request := api.UpdateNetworkServerRequest{
		NetworkServer: &networkServer,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Update(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating network server",
			err.Error(),
		)
		return
	}

	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceNetworkServer) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state
	var state NetworkServer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteNetworkServerRequest{
		Id: state.ID.Value,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Delete(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting network server",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceNetworkServer) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing network server", err.Error())
	}
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), id)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
