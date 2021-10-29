package chirpstack

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkServer struct {
	ID                          types.Int64  `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
	Server                      types.String `tfsdk:"server"`
	CACert                      types.String `tfsdk:"ca_cert"`
	TLSCert                     types.String `tfsdk:"tls_cert"`
	RoutingProfileCACert        types.String `tfsdk:"routing_profil_ca_cert"`
	RoutingProfileTLSCert       types.String `tfsdk:"routing_profil_tls_cert"`
	RoutingProfileTLSKey        types.String `tfsdk:"routing_profil_tls_key"`
	GatewayDiscoveryEnabled     types.Bool   `tfsdk:"gateway_discovery_enabled"`
	GatewayDiscoveryInterval    types.Int64  `tfsdk:"gateway_discovery_interval"`
	GatewayDiscoveryTXFrequency types.Int64  `tfsdk:"gateway_discovery_tx_frequency"`
	GatewayDiscoveryDR          types.Int64  `tfsdk:"gateway_discovery_dr"`
}

type APIKey struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	IsAdmin        types.Bool   `tfsdk:"is_admin"`
	OrganizationID types.Int64  `tfsdk:"organization_id"`
	ApplicationID  types.Int64  `tfsdk:"application_id"`
	Key            types.String `tfsdk:"key"`
}

type Organization struct {
	ID              types.Int64  `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	DisplayName     types.String `tfsdk:"display_name"`
	CanHaveGateways types.Bool   `tfsdk:"can_have_gateways"`
	MaxGatewayCount types.Int64  `tfsdk:"max_gateway_count"`
	MaxDeviceCount  types.Int64  `tfsdk:"max_device_count"`
}
