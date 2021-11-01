package models

import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/brocaar/chirpstack-api/go/v3/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func GatewaySchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"name":        {Type: types.StringType, Required: true},
			"description": {Type: types.StringType, Required: true},
			"location": {
				Type: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"latitude":  types.Float64Type,
						"longitude": types.Float64Type,
						"altitude":  types.Float64Type,
						"source":    types.StringType,
					},
				},
				Required: true,
			},
			"organization_id": {
				Type:     types.Int64Type,
				Required: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
			},
			"discovery_enabled":  {Type: types.BoolType, Optional: true, Computed: true},
			"network_server_id":  {Type: types.Int64Type, Required: true},
			"gateway_profile_id": {Type: types.StringType, Optional: true, Computed: true},
			"tags": {
				Type:     types.MapType{ElemType: types.StringType},
				Optional: true,
				Computed: true,
			},
			"metadata": {
				Type:     types.MapType{ElemType: types.StringType},
				Optional: true,
				Computed: true,
			},
			"service_profile_id": {Type: types.StringType, Optional: true, Computed: true},
		},
	}
}

type Gateway struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	// *common.Location
	Location         types.Object `tfsdk:"location"`
	OrganizationId   types.Int64  `tfsdk:"organization_id"`
	DiscoveryEnabled types.Bool   `tfsdk:"discovery_enabled"`
	NetworkServerId  types.Int64  `tfsdk:"network_server_id"`
	GatewayProfileId types.String `tfsdk:"gateway_profile_id"`
	// TODO: Boards
	// map[string]string
	Tags types.Map `tfsdk:"tags"`
	// map[string]string
	Metadata         types.Map    `tfsdk:"metadata"`
	ServiceProfileId types.String `tfsdk:"service_profile_id"`
}

func GatewayFromApiType(s *api.Gateway) Gateway {
	tags := map[string]attr.Value{}
	for k, v := range s.Tags {
		tags[k] = types.String{Value: v}
	}
	metadata := map[string]attr.Value{}
	for k, v := range s.Metadata {
		metadata[k] = types.String{Value: v}
	}
	return Gateway{
		Id:          types.String{Value: s.Id},
		Name:        types.String{Value: s.Name},
		Description: types.String{Value: s.Description},
		Location: types.Object{
			Attrs: map[string]attr.Value{
				"latitude":  types.Float64{Value: s.Location.Latitude},
				"longitude": types.Float64{Value: s.Location.Longitude},
				"altitude":  types.Float64{Value: s.Location.Altitude},
				"source":    types.String{Value: s.Location.Source.String()},
			},
			AttrTypes: map[string]attr.Type{
				"latitude":  types.Float64Type,
				"longitude": types.Float64Type,
				"altitude":  types.Float64Type,
				"source":    types.StringType,
			},
		},
		OrganizationId:   types.Int64{Value: s.OrganizationId},
		DiscoveryEnabled: types.Bool{Value: s.DiscoveryEnabled},
		NetworkServerId:  types.Int64{Value: s.NetworkServerId},
		GatewayProfileId: types.String{Value: s.GatewayProfileId},
		Tags:             types.Map{Elems: tags, ElemType: types.StringType},
		Metadata:         types.Map{Elems: metadata, ElemType: types.StringType},
		ServiceProfileId: types.String{Value: s.GatewayProfileId},
	}
}

func (s *Gateway) ToApiType(ctx context.Context) api.Gateway {
	tags := map[string]string{}
	for k, v := range s.Tags.Elems {
		value, err := v.ToTerraformValue(ctx)
		if err != nil {
			panic(err)
		}
		tags[k] = value.(types.String).Value
	}
	metadata := map[string]string{}
	for k, v := range s.Metadata.Elems {
		value, err := v.ToTerraformValue(ctx)
		if err != nil {
			panic(err)
		}
		metadata[k] = value.(types.String).Value
	}
	source := s.Location.Attrs["source"].(types.String).Value
	return api.Gateway{
		Id:          s.Id.Value,
		Name:        s.Name.Value,
		Description: s.Description.Value,
		Location: &common.Location{
			Latitude:  s.Location.Attrs["latitude"].(types.Float64).Value,
			Longitude: s.Location.Attrs["longitude"].(types.Float64).Value,
			Altitude:  s.Location.Attrs["altitude"].(types.Float64).Value,
			Source:    common.LocationSource(common.LocationSource_value[source]),
		},
		OrganizationId:   s.OrganizationId.Value,
		DiscoveryEnabled: s.DiscoveryEnabled.Value,
		NetworkServerId:  s.NetworkServerId.Value,
		GatewayProfileId: s.GatewayProfileId.Value,
		Tags:             tags,
		Metadata:         metadata,
		ServiceProfileId: s.ServiceProfileId.Value,
	}
}
