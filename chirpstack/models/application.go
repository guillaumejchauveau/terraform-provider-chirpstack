package models

import (
	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ApplicationSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id":          {Type: types.Int64Type, Computed: true},
			"name":        {Type: types.StringType, Required: true},
			"description": {Type: types.StringType, Required: true},
			"organization_id": {
				Type:     types.Int64Type,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"service_profile_id": {
				Type:     types.StringType,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"payload_codec":          {Type: types.StringType, Optional: true, Computed: true},
			"payload_encoder_script": {Type: types.StringType, Optional: true, Computed: true},
			"payload_decoder_script": {Type: types.StringType, Optional: true, Computed: true},
		},
	}
}

type Application struct {
	Id                   types.Int64  `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	OrganizationId       types.Int64  `tfsdk:"organization_id"`
	ServiceProfileId     types.String `tfsdk:"service_profile_id"`
	PayloadCodec         types.String `tfsdk:"payload_codec"`
	PayloadEncoderScript types.String `tfsdk:"payload_encoder_script"`
	PayloadDecoderScript types.String `tfsdk:"payload_decoder_script"`
}

func ApplicationFromApiType(s *api.Application) Application {
	return Application{
		Id:                   types.Int64{Value: int64(s.Id)},
		Name:                 types.String{Value: s.Name},
		Description:          types.String{Value: s.Description},
		OrganizationId:       types.Int64{Value: s.OrganizationId},
		ServiceProfileId:     types.String{Value: s.ServiceProfileId},
		PayloadCodec:         types.String{Value: s.PayloadCodec},
		PayloadEncoderScript: types.String{Value: s.PayloadEncoderScript},
		PayloadDecoderScript: types.String{Value: s.PayloadDecoderScript},
	}
}

func (s *Application) ToApiType() api.Application {
	return api.Application{
		Id:                   s.Id.Value,
		Name:                 s.Name.Value,
		Description:          s.Description.Value,
		OrganizationId:       s.OrganizationId.Value,
		ServiceProfileId:     s.ServiceProfileId.Value,
		PayloadCodec:         s.PayloadCodec.Value,
		PayloadEncoderScript: s.PayloadEncoderScript.Value,
		PayloadDecoderScript: s.PayloadDecoderScript.Value,
	}
}
