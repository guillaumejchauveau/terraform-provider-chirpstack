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

type ServiceProfile struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	OrganizationId  types.Int64  `tfsdk:"organization_id"`
	NetworkServerId types.Int64  `tfsdk:"network_server_id"`
	UlRate          types.Int64  `tfsdk:"ul_rate"`
	UlBucketSize    types.Int64  `tfsdk:"ul_bucket_size"`
	// See api.RatePolicy
	UlRatePolicy types.Int64 `tfsdk:"ul_rate_policy"`
	DlRate       types.Int64 `tfsdk:"dl_rate"`
	DlBucketSize types.Int64 `tfsdk:"dl_bucket_size"`
	// See api.RatePolicy
	DlRatePolicy           types.Int64 `tfsdk:"dl_rate_policy"`
	AddGwMetadata          types.Bool  `tfsdk:"add_gw_metadata"`
	DevStatusReqFreq       types.Int64 `tfsdk:"dev_status_req_freq"`
	ReportDevStatusBattery types.Bool  `tfsdk:"report_dev_status_battery"`
	ReportDevStatusMargin  types.Bool  `tfsdk:"report_dev_status_margin"`
	DrMin                  types.Int64 `tfsdk:"dr_min"`
	DrMax                  types.Int64 `tfsdk:"dr_max"`
	// []byte
	ChannelMask    types.String `tfsdk:"channel_mask"`
	PrAllowed      types.Bool   `tfsdk:"pr_allowed"`
	HrAllowed      types.Bool   `tfsdk:"hr_allowed"`
	RaAllowed      types.Bool   `tfsdk:"ra_allowed"`
	NwkGeoLoc      types.Bool   `tfsdk:"nwk_geo_loc"`
	TargetPer      types.Int64  `tfsdk:"target_per"`
	MinGwDiversity types.Int64  `tfsdk:"min_gw_diversity"`
	GwsPrivate     types.Bool   `tfsdk:"gws_private"`
}

type DeviceProfile struct {
	ID                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	OrganizationId    types.Int64  `tfsdk:"organization_id"`
	NetworkServerId   types.Int64  `tfsdk:"network_server_id"`
	SupportsClassB    types.Bool   `tfsdk:"supports_class_b"`
	ClassBTimeout     types.Int64  `tfsdk:"class_b_timeout"`
	PingSlotPeriod    types.Int64  `tfsdk:"ping_slot_period"`
	PingSlotDr        types.Int64  `tfsdk:"ping_slot_dr"`
	PingSlotFreq      types.Int64  `tfsdk:"ping_slot_freq"`
	SupportsClassC    types.Bool   `tfsdk:"supports_class_c"`
	ClassCTimeout     types.Int64  `tfsdk:"class_c_timeout"`
	MacVersion        types.String `tfsdk:"mac_version"`
	RegParamsRevision types.String `tfsdk:"reg_params_revision"`
	RxDelay_1         types.Int64  `tfsdk:"rx_delay_1"`
	RxDrOffset_1      types.Int64  `tfsdk:"rx_dr_offset_1"`
	RxDatarate_2      types.Int64  `tfsdk:"rx_datarate_2"`
	RxFreq_2          types.Int64  `tfsdk:"rx_freq_2"`
	// []uint32
	FactoryPresetFreqs   types.Set    `tfsdk:"factory_preset_freqs"`
	MaxEirp              types.Int64  `tfsdk:"max_eirp"`
	MaxDutyCycle         types.Int64  `tfsdk:"max_duty_cycle"`
	SupportsJoin         types.Bool   `tfsdk:"supports_join"`
	RfRegion             types.String `tfsdk:"rf_region"`
	Supports_32BitFCnt   types.Bool   `tfsdk:"supports_32bit_f_cnt"`
	PayloadCodec         types.String `tfsdk:"payload_codec"`
	PayloadEncoderScript types.String `tfsdk:"payload_encoder_script"`
	PayloadDecoderScript types.String `tfsdk:"payload_decoder_script"`
	GeolocBufferTtl      types.Int64  `tfsdk:"geoloc_buffer_ttl"`
	GeolocMinBufferSize  types.Int64  `tfsdk:"geoloc_min_buffer_size"`
	// map[string]string
	Tags           types.Map    `tfsdk:"tags"`
	UplinkInterval Duration     `tfsdk:"uplink_interval"`
	AdrAlgorithmId types.String `tfsdk:"adr_algorithm_id"`
}
