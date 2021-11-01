package models

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func APIKeySchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {Type: types.StringType, Computed: true},
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
			"key": {Type: types.StringType, Computed: true, Sensitive: true},
		},
	}
}

type APIKey struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	IsAdmin        types.Bool   `tfsdk:"is_admin"`
	OrganizationID types.Int64  `tfsdk:"organization_id"`
	ApplicationID  types.Int64  `tfsdk:"application_id"`
	Key            types.String `tfsdk:"key"`
}
