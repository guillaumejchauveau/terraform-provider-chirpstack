package chirpstack

import (
	"context"
	"strconv"
	"strings"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type resourceOrganizationUserType struct{}

func (r resourceOrganizationUserType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"organization_id": {
				Type:     types.Int64Type,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"user_id": {
				Type:     types.Int64Type,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"email": {
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
			},
			"is_device_admin": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"is_gateway_admin": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceOrganizationUserType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceOrganizationUser{
		p: *(p.(*provider)),
	}, nil
}

type resourceOrganizationUser struct {
	p provider
}

// Create a new resource
func (r resourceOrganizationUser) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan OrganizationUser
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	organizationUser := api.OrganizationUser{
		OrganizationId: plan.OrganizationID.Value,
		UserId:         plan.UserID.Value,
		Email:          plan.Email.Value,
		IsAdmin:        plan.IsAdmin.Value,
		IsDeviceAdmin:  plan.IsDeviceAdmin.Value,
		IsGatewayAdmin: plan.IsGatewayAdmin.Value,
	}
	request := api.AddOrganizationUserRequest{
		OrganizationUser: &organizationUser,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.AddUser(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating organization user",
			err.Error(),
		)
		return
	}

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("organization_id"), plan.OrganizationID.Value)
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("user_id"), plan.UserID.Value)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceOrganizationUser) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state OrganizationUser
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetOrganizationUserRequest{
		OrganizationId: state.OrganizationID.Value,
		UserId:         state.UserID.Value,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.GetUser(ctx, &request)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError(
			"Error reading organization user",
			err.Error(),
		)
		return
	}

	state.Email = types.String{Value: response.OrganizationUser.Email}
	state.IsAdmin = types.Bool{Value: response.OrganizationUser.IsAdmin}
	state.IsDeviceAdmin = types.Bool{Value: response.OrganizationUser.IsDeviceAdmin}
	state.IsGatewayAdmin = types.Bool{Value: response.OrganizationUser.IsGatewayAdmin}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceOrganizationUser) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan OrganizationUser
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	organizationUser := api.OrganizationUser{
		OrganizationId: plan.OrganizationID.Value,
		UserId:         plan.UserID.Value,
		Email:          plan.Email.Value,
		IsAdmin:        plan.IsAdmin.Value,
		IsDeviceAdmin:  plan.IsDeviceAdmin.Value,
		IsGatewayAdmin: plan.IsGatewayAdmin.Value,
	}
	request := api.UpdateOrganizationUserRequest{
		OrganizationUser: &organizationUser,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.UpdateUser(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating organization user",
			err.Error(),
		)
		return
	}

	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceOrganizationUser) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state
	var state OrganizationUser
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteOrganizationUserRequest{
		OrganizationId: state.OrganizationID.Value,
		UserId:         state.UserID.Value,
	}

	client := api.NewOrganizationServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.DeleteUser(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting organization user",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceOrganizationUser) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	str := strings.Split(req.ID, ",")
	orgId, err := strconv.ParseInt(str[0], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing organization user", err.Error())
	}
	userId, err := strconv.ParseInt(str[1], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing organization user", err.Error())
	}
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("organization_id"), orgId)
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("user_id"), userId)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
