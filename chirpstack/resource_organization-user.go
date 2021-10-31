package chirpstack

import (
	"context"
	"strconv"
	"strings"
	"terraform-provider-chirpstack/chirpstack/models"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type resourceOrganizationUserType struct{}

func (r resourceOrganizationUserType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return models.OrganizationUserSchema(), nil
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
	var plan models.OrganizationUser
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	organizationUser := plan.ToApiType()
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

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("organization_id"), plan.OrganizationId.Value)
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("user_id"), plan.UserId.Value)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceOrganizationUser) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state models.OrganizationUser
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetOrganizationUserRequest{
		OrganizationId: state.OrganizationId.Value,
		UserId:         state.UserId.Value,
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

	newState := models.OrganizationUserFromApiType(response.OrganizationUser)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceOrganizationUser) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan models.OrganizationUser
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	organizationUser := plan.ToApiType()
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
	var state models.OrganizationUser
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteOrganizationUserRequest{
		OrganizationId: state.OrganizationId.Value,
		UserId:         state.UserId.Value,
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
