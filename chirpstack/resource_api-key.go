package chirpstack

import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type resourceAPIKeyType struct{}

// APIKey Resource schema
func (r resourceAPIKeyType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"is_admin": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"organization_id": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"application_id": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"key": {
				Type:      types.StringType,
				Computed:  true,
				Sensitive: true,
			},
		},
	}, nil
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
	var plan APIKey
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

	client := api.NewInternalServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.CreateAPIKey(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating network server",
			"Could not create network server, unexpected error: "+err.Error(),
		)
		return
	}

	plan.ID = types.String{Value: response.Id}
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
	var state APIKey
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
	conn, err := grpc.Dial(r.p.ConnectionData.Server.Value, dialOpts...)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error checking API key",
			"Could not connect to API: "+err.Error(),
		)
		return
	}

	internalClient := api.NewInternalServiceClient(conn)

	_, err = internalClient.Settings(ctx, &emptypb.Empty{})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error checking API key",
			"Could not read settings, unexpected error: "+err.Error(),
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
	var state APIKey
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteRequest := api.DeleteAPIKeyRequest{
		Id: state.ID.Value,
	}

	internalClient := api.NewInternalServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := internalClient.DeleteAPIKey(ctx, &deleteRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting API key",
			"Could not delete key, unexpected error: "+err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceAPIKey) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
}
