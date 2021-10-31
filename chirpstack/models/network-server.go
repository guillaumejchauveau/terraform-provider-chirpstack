package models

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
)

func NetworkServerSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.Int64Type,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"server": {
				Type:     types.StringType,
				Required: true,
			},
			"ca_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"tls_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"routing_profil_ca_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"routing_profil_tls_cert": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"routing_profil_tls_key": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"gateway_discovery_enabled": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"gateway_discovery_interval": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"gateway_discovery_tx_frequency": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"gateway_discovery_dr": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
		},
	}
}

type NetworkServer struct {
	Id                          types.Int64  `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
	Server                      types.String `tfsdk:"server"`
	CaCert                      types.String `tfsdk:"ca_cert"`
	TlsCert                     types.String `tfsdk:"tls_cert"`
	RoutingProfileCaCert        types.String `tfsdk:"routing_profil_ca_cert"`
	RoutingProfileTlsCert       types.String `tfsdk:"routing_profil_tls_cert"`
	RoutingProfileTlsKey        types.String `tfsdk:"routing_profil_tls_key"`
	GatewayDiscoveryEnabled     types.Bool   `tfsdk:"gateway_discovery_enabled"`
	GatewayDiscoveryInterval    types.Int64  `tfsdk:"gateway_discovery_interval"`
	GatewayDiscoveryTxFrequency types.Int64  `tfsdk:"gateway_discovery_tx_frequency"`
	GatewayDiscoveryDr          types.Int64  `tfsdk:"gateway_discovery_dr"`
}

func NetworkServerFromApiType(s *api.NetworkServer) NetworkServer {
	return NetworkServer{
		Id:                          types.Int64{Value: s.Id},
		Name:                        types.String{Value: s.Name},
		Server:                      types.String{Value: s.Server},
		CaCert:                      types.String{Value: s.CaCert},
		TlsCert:                     types.String{Value: s.TlsCert},
		RoutingProfileCaCert:        types.String{Value: s.RoutingProfileCaCert},
		RoutingProfileTlsCert:       types.String{Value: s.RoutingProfileTlsCert},
		RoutingProfileTlsKey:        types.String{Value: s.RoutingProfileTlsKey},
		GatewayDiscoveryEnabled:     types.Bool{Value: s.GatewayDiscoveryEnabled},
		GatewayDiscoveryInterval:    types.Int64{Value: int64(s.GatewayDiscoveryInterval)},
		GatewayDiscoveryTxFrequency: types.Int64{Value: int64(s.GatewayDiscoveryTxFrequency)},
		GatewayDiscoveryDr:          types.Int64{Value: int64(s.GatewayDiscoveryDr)},
	}
}

func (s *NetworkServer) ToApiType() api.NetworkServer {
	return api.NetworkServer{
		Id:                          s.Id.Value,
		Name:                        s.Name.Value,
		Server:                      s.Server.Value,
		CaCert:                      s.CaCert.Value,
		TlsCert:                     s.TlsCert.Value,
		RoutingProfileCaCert:        s.RoutingProfileCaCert.Value,
		RoutingProfileTlsCert:       s.RoutingProfileTlsCert.Value,
		RoutingProfileTlsKey:        s.RoutingProfileTlsKey.Value,
		GatewayDiscoveryEnabled:     s.GatewayDiscoveryEnabled.Value,
		GatewayDiscoveryInterval:    uint32(s.GatewayDiscoveryInterval.Value),
		GatewayDiscoveryTxFrequency: uint32(s.GatewayDiscoveryTxFrequency.Value),
		GatewayDiscoveryDr:          uint32(s.GatewayDiscoveryDr.Value),
	}
}
