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

type resourceUserType struct{}

func (r resourceUserType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return models.UserSchema(), nil
}

// New resource instance
func (r resourceUserType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceUser{
		p: *(p.(*provider)),
	}, nil
}

type resourceUser struct {
	p provider
}

// Create a new resource
func (r resourceUser) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan models.User
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	user := plan.ToApiType()
	request := api.CreateUserRequest{
		Password: plan.Password.Value,
		User:     &user,
	}

	client := api.NewUserServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Create(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating user",
			err.Error(),
		)
		return
	}

	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), response.Id)
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("password"), plan.Password.Value)

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceUser) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state models.User
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetUserRequest{
		Id: state.Id.Value,
	}

	client := api.NewUserServiceClient(r.p.Conn(ctx))
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
			"Error reading user",
			err.Error(),
		)
		return
	}

	newState := models.UserFromApiType(response.User, state.Password.Value)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceUser) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan models.User
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.User
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id

	client := api.NewUserServiceClient(r.p.Conn(ctx))
	if !state.Password.Equal(plan.Password) {
		request := api.UpdateUserPasswordRequest{
			UserId:   plan.Id.Value,
			Password: plan.Password.Value,
		}
		_, err := client.UpdatePassword(ctx, &request)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating user password",
				err.Error(),
			)
			return
		}

		resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("password"), plan.Password.Value)
		state.Password = plan.Password
	}

	if plan.Equal(state) {
		return
	}

	user := plan.ToApiType()
	request := api.UpdateUserRequest{
		User: &user,
	}

	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Update(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating user",
			err.Error(),
		)
		return
	}

	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceUser) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state
	var state models.User
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteUserRequest{
		Id: state.Id.Value,
	}

	client := api.NewUserServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Delete(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting user",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceUser) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing user", err.Error())
	}
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), id)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
