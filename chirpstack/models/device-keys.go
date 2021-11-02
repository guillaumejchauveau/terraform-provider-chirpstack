package models

import (
	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DeviceKeysSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"dev_eui": {
				Type:     types.StringType,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"nwk_key": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"app_key": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"gen_app_key": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

type DeviceKeys struct {
	DevEui    types.String `tfsdk:"dev_eui"`
	NwkKey    types.String `tfsdk:"nwk_key"`
	AppKey    types.String `tfsdk:"app_key"`
	GenAppKey types.String `tfsdk:"gen_app_key"`
}

func DeviceKeysFromApiType(s *api.DeviceKeys) DeviceKeys {
	return DeviceKeys{
		DevEui:    types.String{Value: s.DevEui},
		NwkKey:    types.String{Value: s.NwkKey},
		AppKey:    types.String{Value: s.AppKey},
		GenAppKey: types.String{Value: s.GenAppKey},
	}
}

func (s *DeviceKeys) ToApiType() api.DeviceKeys {
	return api.DeviceKeys{
		DevEui:    s.DevEui.Value,
		NwkKey:    s.NwkKey.Value,
		AppKey:    s.AppKey.Value,
		GenAppKey: s.GenAppKey.Value,
	}
}
