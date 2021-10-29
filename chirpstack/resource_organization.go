package chirpstack

import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceOrganizationType struct{}

// NetworkServer Resource schema
func (r resourceOrganizationType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
			"display_name": {
				Type:     types.StringType,
				Required: true,
			},
			"can_have_gateways": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"max_gateway_count": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"max_device_count": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceOrganizationType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceOrganization{
		p: *(p.(*provider)),
	}, nil
}

type resourceOrganization struct {
	p provider
}

// Create a new resource
func (r resourceOrganization) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan Organization
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	organization := api.Organization{
		Name:            plan.Name.Value,
		DisplayName:     plan.DisplayName.Value,
		CanHaveGateways: plan.CanHaveGateways.Value,
		MaxGatewayCount: uint32(plan.MaxGatewayCount.Value),
		MaxDeviceCount:  uint32(plan.MaxDeviceCount.Value),
	}
	request := api.CreateOrganizationRequest{
		Organization: &organization,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Create(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating organization",
			"Could not create organization, unexpected error: "+err.Error(),
		)
		return
	}

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), response.Id)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceOrganization) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state Organization
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetOrganizationRequest{
		Id: state.ID.Value,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Get(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading organization",
			"Could not read organization, unexpected error: "+err.Error(),
		)
		return
	}

	state.Name = types.String{Value: response.Organization.Name}
	state.DisplayName = types.String{Value: response.Organization.DisplayName}
	state.CanHaveGateways = types.Bool{Value: response.Organization.CanHaveGateways}
	state.MaxGatewayCount = types.Int64{Value: int64(response.Organization.MaxGatewayCount)}
	state.MaxDeviceCount = types.Int64{Value: int64(response.Organization.MaxDeviceCount)}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceOrganization) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan Organization
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, diags := req.State.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"))
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	organization := api.Organization{
		Id:              id.(types.Int64).Value,
		Name:            plan.Name.Value,
		DisplayName:     plan.DisplayName.Value,
		CanHaveGateways: plan.CanHaveGateways.Value,
		MaxGatewayCount: uint32(plan.MaxGatewayCount.Value),
		MaxDeviceCount:  uint32(plan.MaxDeviceCount.Value),
	}
	request := api.UpdateOrganizationRequest{
		Organization: &organization,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Update(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating organization",
			"Could not update organization, unexpected error: "+err.Error(),
		)
		return
	}

	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceOrganization) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state
	var state Organization
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteOrganizationRequest{
		Id: state.ID.Value,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Delete(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting organization",
			"Could not delete organization, unexpected error: "+err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceOrganization) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
}
