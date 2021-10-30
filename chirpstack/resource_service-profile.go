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

type resourceServiceProfileType struct{}

func (r resourceServiceProfileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
			"organization_id": {
				Type:     types.Int64Type,
				Required: true,
			},
			"network_server_id": {
				Type:     types.Int64Type,
				Required: true,
			},
			"ul_rate": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"ul_bucket_size": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"ul_rate_policy": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dl_rate": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dl_bucket_size": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dl_rate_policy": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"add_gw_metadata": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"dev_status_req_freq": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"report_dev_status_battery": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"report_dev_status_margin": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"dr_min": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dr_max": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"channel_mask": {
				Type:     types.StringType,
				Required: true,
			},
			"pr_allowed": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"hr_allowed": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"ra_allowed": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"nwk_geo_loc": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"target_per": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"min_gw_diversity": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"gws_private": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceServiceProfileType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceServiceProfile{
		p: *(p.(*provider)),
	}, nil
}

type resourceServiceProfile struct {
	p provider
}

// Create a new resource
func (r resourceServiceProfile) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan ServiceProfile
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serviceprofile := api.ServiceProfile{
		Name: plan.Name.Value,
	}
	request := api.CreateServiceProfileRequest{
		ServiceProfile: &serviceprofile,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Create(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating serviceprofile",
			err.Error(),
		)
		return
	}

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), response.Id)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceServiceProfile) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state ServiceProfile
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetServiceProfileRequest{
		Id: state.ID.Value,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Get(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading service profile",
			err.Error(),
		)
		return
	}

	state.Name = types.String{Value: response.ServiceProfile.Name}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceServiceProfile) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan ServiceProfile
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state ServiceProfile
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = state.ID

	serviceprofile := api.ServiceProfile{
		Id:   plan.ID.Value,
		Name: plan.Name.Value,
	}
	request := api.UpdateServiceProfileRequest{
		ServiceProfile: &serviceprofile,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Update(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating service profile",
			err.Error(),
		)
		return
	}

	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceServiceProfile) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state
	var state ServiceProfile
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteServiceProfileRequest{
		Id: state.ID.Value,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Delete(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting service profile",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceServiceProfile) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing service profile", err.Error())
	}
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), id)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
