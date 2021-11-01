package models

import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DeviceSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"dev_eui": {
				Type:     types.StringType,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"name":               {Type: types.StringType, Required: true},
			"application_id":     {Type: types.Int64Type, Required: true},
			"description":        {Type: types.StringType, Required: true},
			"device_profile_id":  {Type: types.StringType, Required: true},
			"skip_f_cnt_check":   {Type: types.BoolType, Optional: true, Computed: true},
			"reference_altitude": {Type: types.Float64Type, Optional: true, Computed: true},
			"variables": {
				Type:     types.MapType{ElemType: types.StringType},
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     types.MapType{ElemType: types.StringType},
				Optional: true,
				Computed: true,
			},
			"is_disabled": {Type: types.BoolType, Optional: true, Computed: true},
		},
	}
}

type Device struct {
	DevEui            types.String  `tfsdk:"dev_eui"`
	Name              types.String  `tfsdk:"name"`
	ApplicationId     types.Int64   `tfsdk:"application_id"`
	Description       types.String  `tfsdk:"description"`
	DeviceProfileId   types.String  `tfsdk:"device_profile_id"`
	SkipFCntCheck     types.Bool    `tfsdk:"skip_f_cnt_check"`
	ReferenceAltitude types.Float64 `tfsdk:"reference_altitude"`
	// map[string]string
	Variables types.Map `tfsdk:"variables"`
	// map[string]string
	Tags       types.Map  `tfsdk:"tags"`
	IsDisabled types.Bool `tfsdk:"is_disabled"`
}

func DeviceFromApiType(s *api.Device) Device {
	variables := map[string]attr.Value{}
	for k, v := range s.Variables {
		variables[k] = types.String{Value: v}
	}
	tags := map[string]attr.Value{}
	for k, v := range s.Tags {
		tags[k] = types.String{Value: v}
	}
	return Device{
		DevEui:            types.String{Value: s.DevEui},
		Name:              types.String{Value: s.Name},
		ApplicationId:     types.Int64{Value: s.ApplicationId},
		Description:       types.String{Value: s.Description},
		DeviceProfileId:   types.String{Value: s.DeviceProfileId},
		SkipFCntCheck:     types.Bool{Value: s.SkipFCntCheck},
		ReferenceAltitude: types.Float64{Value: s.ReferenceAltitude},
		Variables:         types.Map{Elems: variables, ElemType: types.StringType},
		Tags:              types.Map{Elems: tags, ElemType: types.StringType},
		IsDisabled:        types.Bool{Value: s.IsDisabled},
	}
}

func (s *Device) ToApiType(ctx context.Context) api.Device {
	variables := map[string]string{}
	for k, v := range s.Variables.Elems {
		value, err := v.ToTerraformValue(ctx)
		if err != nil {
			panic(err)
		}
		variables[k] = value.(types.String).Value
	}
	tags := map[string]string{}
	for k, v := range s.Tags.Elems {
		value, err := v.ToTerraformValue(ctx)
		if err != nil {
			panic(err)
		}
		tags[k] = value.(types.String).Value
	}
	return api.Device{
		DevEui:            s.DevEui.Value,
		Name:              s.Name.Value,
		ApplicationId:     s.ApplicationId.Value,
		Description:       s.Description.Value,
		DeviceProfileId:   s.DeviceProfileId.Value,
		SkipFCntCheck:     s.SkipFCntCheck.Value,
		ReferenceAltitude: s.ReferenceAltitude.Value,
		Variables:         variables,
		Tags:              tags,
		IsDisabled:        s.IsDisabled.Value,
	}
}
