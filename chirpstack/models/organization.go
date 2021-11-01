package models

import (
	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func OrganizationSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id":                {Type: types.Int64Type, Computed: true},
			"name":              {Type: types.StringType, Required: true},
			"display_name":      {Type: types.StringType, Required: true},
			"can_have_gateways": {Type: types.BoolType, Optional: true, Computed: true},
			"max_gateway_count": {Type: types.Int64Type, Optional: true, Computed: true},
			"max_device_count":  {Type: types.Int64Type, Optional: true, Computed: true},
		},
	}
}

type Organization struct {
	Id              types.Int64  `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	DisplayName     types.String `tfsdk:"display_name"`
	CanHaveGateways types.Bool   `tfsdk:"can_have_gateways"`
	MaxGatewayCount types.Int64  `tfsdk:"max_gateway_count"`
	MaxDeviceCount  types.Int64  `tfsdk:"max_device_count"`
}

func OrganizationFromApiType(s *api.Organization) Organization {
	return Organization{
		Id:              types.Int64{Value: int64(s.Id)},
		Name:            types.String{Value: s.Name},
		DisplayName:     types.String{Value: s.DisplayName},
		CanHaveGateways: types.Bool{Value: s.CanHaveGateways},
		MaxGatewayCount: types.Int64{Value: int64(s.MaxGatewayCount)},
		MaxDeviceCount:  types.Int64{Value: int64(s.MaxDeviceCount)},
	}
}

func (s *Organization) ToApiType() api.Organization {
	return api.Organization{
		Id:              s.Id.Value,
		Name:            s.Name.Value,
		DisplayName:     s.DisplayName.Value,
		CanHaveGateways: s.CanHaveGateways.Value,
		MaxGatewayCount: uint32(s.MaxGatewayCount.Value),
		MaxDeviceCount:  uint32(s.MaxDeviceCount.Value),
	}
}
