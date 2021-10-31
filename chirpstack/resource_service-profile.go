package chirpstack

import (
	"context"
	"terraform-provider-chirpstack/chirpstack/models"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type resourceServiceProfileType struct{}

func (r resourceServiceProfileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return models.ServiceProfileSchema(), nil
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
	var plan models.ServiceProfile
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serviceprofile := plan.ToApiType()
	request := api.CreateServiceProfileRequest{
		ServiceProfile: &serviceprofile,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
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

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceServiceProfile) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state models.ServiceProfile
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetServiceProfileRequest{
		Id: state.Id.Value,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Get(ctx, &request)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError(
			"Error reading service profile",
			err.Error(),
		)
		return
	}

	newState := models.ServiceProfileFromApiType(response.ServiceProfile)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceServiceProfile) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan models.ServiceProfile
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.ServiceProfile
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id

	serviceprofile := plan.ToApiType()
	request := api.UpdateServiceProfileRequest{
		ServiceProfile: &serviceprofile,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
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
	var state models.ServiceProfile
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteServiceProfileRequest{
		Id: state.Id.Value,
	}

	client := api.NewServiceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
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
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req.ID)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
