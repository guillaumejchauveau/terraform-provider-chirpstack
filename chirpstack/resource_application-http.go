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

type resourceApplicationHttpType struct{}

func (r resourceApplicationHttpType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return models.ApplicationHttpSchema(), nil
}

// New resource instance
func (r resourceApplicationHttpType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceApplicationHttp{
		p: *(p.(*provider)),
	}, nil
}

type resourceApplicationHttp struct {
	p provider
}

// Create a new resource
func (r resourceApplicationHttp) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan models.ApplicationHttp
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	applicationHttp := plan.ToApiType(ctx)
	request := api.CreateHTTPIntegrationRequest{
		Integration: &applicationHttp,
	}

	client := api.NewApplicationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.CreateHTTPIntegration(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating applicationhttp",
			err.Error(),
		)
		return
	}

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("application_id"), plan.ApplicationId.Value)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceApplicationHttp) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state models.ApplicationHttp
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetHTTPIntegrationRequest{
		ApplicationId: state.ApplicationId.Value,
	}

	client := api.NewApplicationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.GetHTTPIntegration(ctx, &request)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError(
			"Error reading applicationhttp",
			err.Error(),
		)
		return
	}

	newState := models.ApplicationHttpFromApiType(response.Integration)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceApplicationHttp) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan models.ApplicationHttp
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	applicationHttp := plan.ToApiType(ctx)
	request := api.UpdateHTTPIntegrationRequest{
		Integration: &applicationHttp,
	}

	client := api.NewApplicationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.UpdateHTTPIntegration(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating applicationhttp",
			err.Error(),
		)
		return
	}
	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceApplicationHttp) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state.
	var state models.ApplicationHttp
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteRequest := api.DeleteHTTPIntegrationRequest{
		ApplicationId: state.ApplicationId.Value,
	}

	internalClient := api.NewApplicationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := internalClient.DeleteHTTPIntegration(ctx, &deleteRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting applicationhttp",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceApplicationHttp) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing applicationhttp", err.Error())
	}
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("application_id"), id)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
