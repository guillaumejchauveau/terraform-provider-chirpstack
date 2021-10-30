package chirpstack

import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceUserType struct{}

// NetworkServer Resource schema
func (r resourceUserType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.Int64Type,
				Computed: true,
			},
			"email": {
				Type:     types.StringType,
				Required: true,
			},
			"password": {
				Type:      types.StringType,
				Required:  true,
				Sensitive: true,
			},
			"is_active": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"is_admin": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"note": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"session_ttl": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
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
	var plan User
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	user := api.User{
		Email:      plan.Email.Value,
		IsActive:   plan.IsActive.Value,
		IsAdmin:    plan.IsAdmin.Value,
		Note:       plan.Note.Value,
		SessionTtl: int32(plan.SessionTTL.Value),
	}
	request := api.CreateUserRequest{
		Password: plan.Password.Value,
		User:     &user,
	}

	client := api.NewUserServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
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
	var state User
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetUserRequest{
		Id: state.ID.Value,
	}

	client := api.NewUserServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Get(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading user",
			err.Error(),
		)
		return
	}

	state.Email = types.String{Value: response.User.Email}
	state.IsActive = types.Bool{Value: response.User.IsActive}
	state.IsAdmin = types.Bool{Value: response.User.IsAdmin}
	state.Note = types.String{Value: response.User.Note}
	state.SessionTTL = types.Int64{Value: int64(response.User.SessionTtl)}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceUser) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan User
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state User
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = state.ID

	client := api.NewUserServiceClient(r.p.Conn())
	if !state.Password.Equal(plan.Password) {
		request := api.UpdateUserPasswordRequest{
			UserId:   plan.ID.Value,
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

	user := api.User{
		Id:         plan.ID.Value,
		Email:      plan.Email.Value,
		IsActive:   plan.IsActive.Value,
		IsAdmin:    plan.IsAdmin.Value,
		Note:       plan.Note.Value,
		SessionTtl: int32(plan.SessionTTL.Value),
	}
	request := api.UpdateUserRequest{
		User: &user,
	}

	resp.Diagnostics.Append(r.p.Diagnostics...)
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
	var state User
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteUserRequest{
		Id: state.ID.Value,
	}

	client := api.NewUserServiceClient(r.p.Conn())
	resp.Diagnostics.Append(r.p.Diagnostics...)
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
}
