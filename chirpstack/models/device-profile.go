package models

import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DeviceProfileSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id":                  {Type: types.StringType, Computed: true},
			"name":                {Type: types.StringType, Required: true},
			"organization_id":     {Type: types.Int64Type, Required: true},
			"network_server_id":   {Type: types.Int64Type, Required: true},
			"supports_class_b":    {Type: types.BoolType, Optional: true, Computed: true},
			"class_b_timeout":     {Type: types.Int64Type, Optional: true, Computed: true},
			"ping_slot_period":    {Type: types.Int64Type, Optional: true, Computed: true},
			"ping_slot_dr":        {Type: types.Int64Type, Optional: true, Computed: true},
			"ping_slot_freq":      {Type: types.Int64Type, Optional: true, Computed: true},
			"supports_class_c":    {Type: types.BoolType, Optional: true, Computed: true},
			"class_c_timeout":     {Type: types.Int64Type, Optional: true, Computed: true},
			"mac_version":         {Type: types.StringType, Optional: true, Computed: true},
			"reg_params_revision": {Type: types.StringType, Optional: true, Computed: true},
			"rx_delay_1":          {Type: types.Int64Type, Optional: true, Computed: true},
			"rx_dr_offset_1":      {Type: types.Int64Type, Optional: true, Computed: true},
			"rx_datarate_2":       {Type: types.Int64Type, Optional: true, Computed: true},
			"rx_freq_2":           {Type: types.Int64Type, Optional: true, Computed: true},
			"factory_preset_freqs": {
				Type:     types.SetType{ElemType: types.Int64Type},
				Optional: true,
				Computed: true,
			},
			"max_eirp":               {Type: types.Int64Type, Optional: true, Computed: true},
			"max_duty_cycle":         {Type: types.Int64Type, Optional: true, Computed: true},
			"supports_join":          {Type: types.BoolType, Optional: true, Computed: true},
			"rf_region":              {Type: types.StringType, Optional: true, Computed: true},
			"supports_32bit_f_cnt":   {Type: types.BoolType, Optional: true, Computed: true},
			"payload_codec":          {Type: types.StringType, Optional: true, Computed: true},
			"payload_encoder_script": {Type: types.StringType, Optional: true, Computed: true},
			"payload_decoder_script": {Type: types.StringType, Optional: true, Computed: true},
			"geoloc_buffer_ttl":      {Type: types.Int64Type, Optional: true, Computed: true},
			"geoloc_min_buffer_size": {Type: types.Int64Type, Optional: true, Computed: true},
			"tags": {
				Type:     types.MapType{ElemType: types.StringType},
				Optional: true,
				Computed: true,
			},
			"uplink_interval":  {Type: DurationType{}, Optional: true, Computed: true},
			"adr_algorithm_id": {Type: types.StringType, Optional: true, Computed: true},
		},
	}
}

type DeviceProfile struct {
	Id                types.String `tfsdk:"id"`
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

func DeviceProfileFromApiType(s *api.DeviceProfile) DeviceProfile {
	var factoryPresetFreqs []attr.Value
	for _, v := range s.FactoryPresetFreqs {
		factoryPresetFreqs = append(factoryPresetFreqs, types.Int64{Value: int64(v)})
	}
	tags := map[string]attr.Value{}
	for k, v := range s.Tags {
		tags[k] = types.String{Value: v}
	}

	return DeviceProfile{
		Id:                   types.String{Value: s.Id},
		Name:                 types.String{Value: s.Name},
		OrganizationId:       types.Int64{Value: s.OrganizationId},
		NetworkServerId:      types.Int64{Value: s.NetworkServerId},
		SupportsClassB:       types.Bool{Value: s.SupportsClassB},
		ClassBTimeout:        types.Int64{Value: int64(s.ClassBTimeout)},
		PingSlotPeriod:       types.Int64{Value: int64(s.PingSlotPeriod)},
		PingSlotDr:           types.Int64{Value: int64(s.PingSlotDr)},
		PingSlotFreq:         types.Int64{Value: int64(s.PingSlotFreq)},
		SupportsClassC:       types.Bool{Value: s.SupportsClassC},
		ClassCTimeout:        types.Int64{Value: int64(s.ClassCTimeout)},
		MacVersion:           types.String{Value: s.MacVersion},
		RegParamsRevision:    types.String{Value: s.RegParamsRevision},
		RxDelay_1:            types.Int64{Value: int64(s.RxDelay_1)},
		RxDrOffset_1:         types.Int64{Value: int64(s.RxDrOffset_1)},
		RxDatarate_2:         types.Int64{Value: int64(s.RxDatarate_2)},
		RxFreq_2:             types.Int64{Value: int64(s.RxFreq_2)},
		FactoryPresetFreqs:   types.Set{Elems: factoryPresetFreqs, ElemType: types.Int64Type},
		MaxEirp:              types.Int64{Value: int64(s.MaxEirp)},
		MaxDutyCycle:         types.Int64{Value: int64(s.MaxDutyCycle)},
		SupportsJoin:         types.Bool{Value: s.SupportsJoin},
		RfRegion:             types.String{Value: s.RfRegion},
		Supports_32BitFCnt:   types.Bool{Value: s.Supports_32BitFCnt},
		PayloadCodec:         types.String{Value: s.PayloadCodec},
		PayloadEncoderScript: types.String{Value: s.PayloadEncoderScript},
		PayloadDecoderScript: types.String{Value: s.PayloadDecoderScript},
		GeolocBufferTtl:      types.Int64{Value: int64(s.GeolocBufferTtl)},
		GeolocMinBufferSize:  types.Int64{Value: int64(s.GeolocMinBufferSize)},
		Tags:                 types.Map{Elems: tags, ElemType: types.StringType},
		UplinkInterval:       Duration{Value: s.UplinkInterval},
		AdrAlgorithmId:       types.String{Value: s.AdrAlgorithmId},
	}
}

func (s *DeviceProfile) ToApiType(ctx context.Context) api.DeviceProfile {
	var factoryPresetFreqs []uint32
	for _, v := range s.FactoryPresetFreqs.Elems {
		value, err := v.ToTerraformValue(ctx)
		if err != nil {
			panic(err)
		}
		factoryPresetFreqs = append(factoryPresetFreqs, uint32(value.(types.Int64).Value))
	}
	tags := map[string]string{}
	for k, v := range s.Tags.Elems {
		value, err := v.ToTerraformValue(ctx)
		if err != nil {
			panic(err)
		}
		tags[k] = value.(types.String).Value
	}

	return api.DeviceProfile{
		Id:                   s.Id.Value,
		Name:                 s.Name.Value,
		OrganizationId:       s.OrganizationId.Value,
		NetworkServerId:      s.NetworkServerId.Value,
		SupportsClassB:       s.SupportsClassB.Value,
		ClassBTimeout:        uint32(s.ClassBTimeout.Value),
		PingSlotPeriod:       uint32(s.PingSlotPeriod.Value),
		PingSlotDr:           uint32(s.PingSlotDr.Value),
		PingSlotFreq:         uint32(s.PingSlotFreq.Value),
		SupportsClassC:       s.SupportsClassC.Value,
		ClassCTimeout:        uint32(s.ClassCTimeout.Value),
		MacVersion:           s.MacVersion.Value,
		RegParamsRevision:    s.RegParamsRevision.Value,
		RxDelay_1:            uint32(s.RxDelay_1.Value),
		RxDrOffset_1:         uint32(s.RxDrOffset_1.Value),
		RxDatarate_2:         uint32(s.RxDatarate_2.Value),
		RxFreq_2:             uint32(s.RxFreq_2.Value),
		FactoryPresetFreqs:   factoryPresetFreqs,
		MaxEirp:              uint32(s.MaxEirp.Value),
		MaxDutyCycle:         uint32(s.MaxDutyCycle.Value),
		SupportsJoin:         s.SupportsJoin.Value,
		RfRegion:             s.RfRegion.Value,
		Supports_32BitFCnt:   s.Supports_32BitFCnt.Value,
		PayloadCodec:         s.PayloadCodec.Value,
		PayloadEncoderScript: s.PayloadEncoderScript.Value,
		PayloadDecoderScript: s.PayloadDecoderScript.Value,
		GeolocBufferTtl:      uint32(s.GeolocBufferTtl.Value),
		GeolocMinBufferSize:  uint32(s.GeolocMinBufferSize.Value),
		Tags:                 tags,
		UplinkInterval:       s.UplinkInterval.Value,
		AdrAlgorithmId:       s.AdrAlgorithmId.Value,
	}
}
