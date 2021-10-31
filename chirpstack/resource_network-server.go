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

type resourceNetworkServerType struct{}

func (r resourceNetworkServerType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return models.NetworkServerSchema(), nil
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
	var plan models.NetworkServer
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	networkServer := plan.ToApiType()
	request := api.CreateNetworkServerRequest{
		NetworkServer: &networkServer,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
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
	var state models.NetworkServer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetNetworkServerRequest{
		Id: state.Id.Value,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
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
			"Error reading network server",
			err.Error(),
		)
		return
	}

	newState := models.NetworkServerFromApiType(response.NetworkServer)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceNetworkServer) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan models.NetworkServer
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.NetworkServer
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id

	networkServer := plan.ToApiType()
	request := api.UpdateNetworkServerRequest{
		NetworkServer: &networkServer,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
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
	var state models.NetworkServer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteNetworkServerRequest{
		Id: state.Id.Value,
	}

	client := api.NewNetworkServerServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
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
