package models

import (
	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func OrganizationUserSchema() tfsdk.Schema {
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
	}
}

type OrganizationUser struct {
	OrganizationId types.Int64  `tfsdk:"organization_id"`
	UserId         types.Int64  `tfsdk:"user_id"`
	Email          types.String `tfsdk:"email"`
	IsAdmin        types.Bool   `tfsdk:"is_admin"`
	IsDeviceAdmin  types.Bool   `tfsdk:"is_device_admin"`
	IsGatewayAdmin types.Bool   `tfsdk:"is_gateway_admin"`
}

func OrganizationUserFromApiType(s *api.OrganizationUser) OrganizationUser {
	return OrganizationUser{
		OrganizationId: types.Int64{Value: int64(s.OrganizationId)},
		UserId:         types.Int64{Value: int64(s.UserId)},
		Email:          types.String{Value: s.Email},
		IsAdmin:        types.Bool{Value: s.IsAdmin},
		IsDeviceAdmin:  types.Bool{Value: s.IsDeviceAdmin},
		IsGatewayAdmin: types.Bool{Value: s.IsGatewayAdmin},
	}
}

func (s *OrganizationUser) ToApiType() api.OrganizationUser {
	return api.OrganizationUser{
		OrganizationId: s.OrganizationId.Value,
		UserId:         s.UserId.Value,
		Email:          s.Email.Value,
		IsAdmin:        s.IsAdmin.Value,
		IsDeviceAdmin:  s.IsDeviceAdmin.Value,
		IsGatewayAdmin: s.IsGatewayAdmin.Value,
	}
}
