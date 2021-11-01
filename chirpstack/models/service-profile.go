package models

import (
	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ServiceProfileSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id":                  {Type: types.StringType, Computed: true},
			"name":                {Type: types.StringType, Required: true},
			"organization_id":     {Type: types.Int64Type, Required: true},
			"network_server_id":   {Type: types.Int64Type, Required: true},
			"ul_rate":             {Type: types.Int64Type, Optional: true, Computed: true},
			"ul_bucket_size":      {Type: types.Int64Type, Optional: true, Computed: true},
			"ul_rate_policy":      {Type: types.Int64Type, Optional: true, Computed: true},
			"dl_rate":             {Type: types.Int64Type, Optional: true, Computed: true},
			"dl_bucket_size":      {Type: types.Int64Type, Optional: true, Computed: true},
			"dl_rate_policy":      {Type: types.Int64Type, Optional: true, Computed: true},
			"add_gw_metadata":     {Type: types.BoolType, Optional: true, Computed: true},
			"dev_status_req_freq": {Type: types.Int64Type, Optional: true, Computed: true},
			"report_dev_status_battery": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true},
			"report_dev_status_margin": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true},
			"dr_min":           {Type: types.Int64Type, Optional: true, Computed: true},
			"dr_max":           {Type: types.Int64Type, Optional: true, Computed: true},
			"channel_mask":     {Type: types.Int64Type, Optional: true, Computed: true},
			"pr_allowed":       {Type: types.BoolType, Optional: true, Computed: true},
			"hr_allowed":       {Type: types.BoolType, Optional: true, Computed: true},
			"ra_allowed":       {Type: types.BoolType, Optional: true, Computed: true},
			"nwk_geo_loc":      {Type: types.BoolType, Optional: true, Computed: true},
			"target_per":       {Type: types.Int64Type, Optional: true, Computed: true},
			"min_gw_diversity": {Type: types.Int64Type, Optional: true, Computed: true},
			"gws_private":      {Type: types.BoolType, Optional: true, Computed: true},
		},
	}
}

type ServiceProfile struct {
	Id              types.String `tfsdk:"id"`
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
	ChannelMask    types.Int64 `tfsdk:"channel_mask"`
	PrAllowed      types.Bool  `tfsdk:"pr_allowed"`
	HrAllowed      types.Bool  `tfsdk:"hr_allowed"`
	RaAllowed      types.Bool  `tfsdk:"ra_allowed"`
	NwkGeoLoc      types.Bool  `tfsdk:"nwk_geo_loc"`
	TargetPer      types.Int64 `tfsdk:"target_per"`
	MinGwDiversity types.Int64 `tfsdk:"min_gw_diversity"`
	GwsPrivate     types.Bool  `tfsdk:"gws_private"`
}

func ServiceProfileFromApiType(s *api.ServiceProfile) ServiceProfile {
	return ServiceProfile{
		Id:                     types.String{Value: s.Id},
		Name:                   types.String{Value: s.Name},
		OrganizationId:         types.Int64{Value: int64(s.OrganizationId)},
		NetworkServerId:        types.Int64{Value: int64(s.NetworkServerId)},
		UlRate:                 types.Int64{Value: int64(s.UlRate)},
		UlBucketSize:           types.Int64{Value: int64(s.UlBucketSize)},
		UlRatePolicy:           types.Int64{Value: int64(s.UlRatePolicy)},
		DlRate:                 types.Int64{Value: int64(s.DlRate)},
		DlBucketSize:           types.Int64{Value: int64(s.DlBucketSize)},
		DlRatePolicy:           types.Int64{Value: int64(s.DlRatePolicy)},
		AddGwMetadata:          types.Bool{Value: s.AddGwMetadata},
		DevStatusReqFreq:       types.Int64{Value: int64(s.DevStatusReqFreq)},
		ReportDevStatusBattery: types.Bool{Value: s.ReportDevStatusBattery},
		ReportDevStatusMargin:  types.Bool{Value: s.ReportDevStatusMargin},
		DrMin:                  types.Int64{Value: int64(s.DrMin)},
		DrMax:                  types.Int64{Value: int64(s.DrMax)},
		ChannelMask:            types.Int64{Null: true}, // TODO
		PrAllowed:              types.Bool{Value: s.PrAllowed},
		HrAllowed:              types.Bool{Value: s.HrAllowed},
		RaAllowed:              types.Bool{Value: s.RaAllowed},
		NwkGeoLoc:              types.Bool{Value: s.NwkGeoLoc},
		TargetPer:              types.Int64{Value: int64(s.TargetPer)},
		MinGwDiversity:         types.Int64{Value: int64(s.MinGwDiversity)},
		GwsPrivate:             types.Bool{Value: s.GwsPrivate},
	}
}

func (s *ServiceProfile) ToApiType() api.ServiceProfile {
	return api.ServiceProfile{
		Name:                   s.Name.Value,
		OrganizationId:         s.OrganizationId.Value,
		NetworkServerId:        s.NetworkServerId.Value,
		UlRate:                 uint32(s.UlRate.Value),
		UlBucketSize:           uint32(s.UlBucketSize.Value),
		UlRatePolicy:           api.RatePolicy(s.UlRatePolicy.Value),
		DlRate:                 uint32(s.DlRate.Value),
		DlBucketSize:           uint32(s.DlBucketSize.Value),
		DlRatePolicy:           api.RatePolicy(s.DlRatePolicy.Value),
		AddGwMetadata:          s.AddGwMetadata.Value,
		DevStatusReqFreq:       uint32(s.DevStatusReqFreq.Value),
		ReportDevStatusBattery: s.ReportDevStatusBattery.Value,
		ReportDevStatusMargin:  s.ReportDevStatusMargin.Value,
		DrMin:                  uint32(s.DrMin.Value),
		DrMax:                  uint32(s.DrMax.Value),
		ChannelMask:            nil, // TODO
		PrAllowed:              s.PrAllowed.Value,
		HrAllowed:              s.HrAllowed.Value,
		RaAllowed:              s.RaAllowed.Value,
		NwkGeoLoc:              s.NwkGeoLoc.Value,
		TargetPer:              uint32(s.TargetPer.Value),
		MinGwDiversity:         uint32(s.MinGwDiversity.Value),
		GwsPrivate:             s.GwsPrivate.Value,
	}
}
