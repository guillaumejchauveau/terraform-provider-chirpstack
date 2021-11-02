package chirpstack

import (
	"context"
	"terraform-provider-chirpstack/chirpstack/models"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type resourceAPIKeyType struct{}

func (r resourceAPIKeyType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return models.APIKeySchema(), nil
}

// New resource instance
func (r resourceAPIKeyType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceAPIKey{
		p: *(p.(*provider)),
	}, nil
}

type resourceAPIKey struct {
	p provider
}

// Create a new resource
func (r resourceAPIKey) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan models.APIKey
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := api.APIKey{
		Name:           plan.Name.Value,
		IsAdmin:        plan.IsAdmin.Value,
		OrganizationId: plan.OrganizationID.Value,
		ApplicationId:  plan.ApplicationID.Value,
	}
	request := api.CreateAPIKeyRequest{
		ApiKey: &apiKey,
	}

	client := api.NewInternalServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.CreateAPIKey(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating API key",
			err.Error(),
		)
		return
	}

	plan.Id = types.String{Value: response.Id}
	plan.Key = types.String{Value: response.JwtToken}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information
func (r resourceAPIKey) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state models.APIKey
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.p.ConnectionData.Server.Unknown || r.p.ConnectionData.Server.Null {
		resp.Diagnostics.AddError(
			"Error checking API key",
			"Server is null or unknown",
		)
		return
	}

	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(APIToken(state.Key.Value)),
		grpc.WithInsecure(), // remove this when using TLS
	}
	_, err := grpc.Dial(r.p.ConnectionData.Server.Value, dialOpts...)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.Unauthenticated {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError(
			"Error checking API key",
			"Could not connect to API: "+err.Error(),
		)
		return
	}
}

// Update resource
func (r resourceAPIKey) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

// Delete resource
func (r resourceAPIKey) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state.
	var state models.APIKey
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteRequest := api.DeleteAPIKeyRequest{
		Id: state.Id.Value,
	}

	internalClient := api.NewInternalServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := internalClient.DeleteAPIKey(ctx, &deleteRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting API key",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceAPIKey) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	resp.Diagnostics.AddError("Cannot import an API key", "The key can only be known on its creation.")
}
