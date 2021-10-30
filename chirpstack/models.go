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

type User struct {
	ID         types.Int64  `tfsdk:"id"`
	Email      types.String `tfsdk:"email"`
	Password   types.String `tfsdk:"password"`
	IsActive   types.Bool   `tfsdk:"is_active"`
	IsAdmin    types.Bool   `tfsdk:"is_admin"`
	Note       types.String `tfsdk:"note"`
	SessionTTL types.Int64  `tfsdk:"session_ttl"`
}

func (u *User) Equal(o User) bool {
	return u.Email.Value == o.Email.Value &&
		u.Password.Value == o.Password.Value &&
		u.IsActive.Value == o.IsActive.Value &&
		u.IsAdmin.Value == o.IsAdmin.Value &&
		u.Note.Value == o.Note.Value &&
		u.SessionTTL.Value == o.SessionTTL.Value
}

type Organization struct {
	ID              types.Int64  `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	DisplayName     types.String `tfsdk:"display_name"`
	CanHaveGateways types.Bool   `tfsdk:"can_have_gateways"`
	MaxGatewayCount types.Int64  `tfsdk:"max_gateway_count"`
	MaxDeviceCount  types.Int64  `tfsdk:"max_device_count"`
}

type OrganizationUser struct {
	OrganizationID types.Int64  `tfsdk:"organization_id"`
	UserID         types.Int64  `tfsdk:"user_id"`
	Email          types.String `tfsdk:"email"`
	IsAdmin        types.Bool   `tfsdk:"is_admin"`
	IsDeviceAdmin  types.Bool   `tfsdk:"is_device_admin"`
	IsGatewayAdmin types.Bool   `tfsdk:"is_gateway_admin"`
}
