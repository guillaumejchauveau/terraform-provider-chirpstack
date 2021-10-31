package chirpstack

import (
	"context"
	"strconv"
	"terraform-provider-chirpstack/chirpstack/models"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type resourceOrganizationType struct{}

func (r resourceOrganizationType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return models.OrganizationSchema(), nil
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
	var plan models.Organization
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	organization := plan.ToApiType()
	request := api.CreateOrganizationRequest{
		Organization: &organization,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Create(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating organization",
			err.Error(),
		)
		return
	}

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), response.Id)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceOrganization) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state models.Organization
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetOrganizationRequest{
		Id: state.Id.Value,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
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
			"Error reading organization",
			err.Error(),
		)
		return
	}

	newState := models.OrganizationFromApiType(response.Organization)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceOrganization) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan models.Organization
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.Organization
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id

	organization := plan.ToApiType()
	request := api.UpdateOrganizationRequest{
		Organization: &organization,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Update(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating organization",
			err.Error(),
		)
		return
	}

	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceOrganization) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state
	var state models.Organization
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteOrganizationRequest{
		Id: state.Id.Value,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Delete(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting organization",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceOrganization) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing organization", err.Error())
	}
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), id)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
