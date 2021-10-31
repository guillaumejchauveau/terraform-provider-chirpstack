package chirpstack

/*
import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type resourceDeviceProfileType struct{}

func (r resourceDeviceProfileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"organization_id": {
				Type:     types.Int64Type,
				Required: true,
			},
			"network_server_id": {
				Type:     types.Int64Type,
				Required: true,
			},
			"ul_rate": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"ul_bucket_size": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"ul_rate_policy": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dl_rate": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dl_bucket_size": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dl_rate_policy": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"add_gw_metadata": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"dev_status_req_freq": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"report_dev_status_battery": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"report_dev_status_margin": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"dr_min": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"dr_max": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"channel_mask": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"pr_allowed": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"hr_allowed": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"ra_allowed": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"nwk_geo_loc": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
			"target_per": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"min_gw_diversity": {
				Type:     types.Int64Type,
				Optional: true,
				Computed: true,
			},
			"gws_private": {
				Type:     types.BoolType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceDeviceProfileType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceDeviceProfile{
		p: *(p.(*provider)),
	}, nil
}

type resourceDeviceProfile struct {
	p provider
}

// Create a new resource
func (r resourceDeviceProfile) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	// Retrieve values from plan
	var plan DeviceProfile
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceprofile := api.DeviceProfile{
		Name:                   plan.Name.Value,
		OrganizationId:         plan.OrganizationId.Value,
		NetworkServerId:        plan.NetworkServerId.Value,
		UlRate:                 uint32(plan.UlRate.Value),
		UlBucketSize:           uint32(plan.UlBucketSize.Value),
		UlRatePolicy:           api.RatePolicy(plan.UlRatePolicy.Value),
		DlRate:                 uint32(plan.DlRate.Value),
		DlBucketSize:           uint32(plan.DlBucketSize.Value),
		DlRatePolicy:           api.RatePolicy(plan.DlRatePolicy.Value),
		AddGwMetadata:          plan.AddGwMetadata.Value,
		DevStatusReqFreq:       uint32(plan.DevStatusReqFreq.Value),
		ReportDevStatusBattery: plan.ReportDevStatusBattery.Value,
		ReportDevStatusMargin:  plan.ReportDevStatusMargin.Value,
		DrMin:                  uint32(plan.DrMin.Value),
		DrMax:                  uint32(plan.DrMax.Value),
		ChannelMask:            nil, // TODO
		PrAllowed:              plan.PrAllowed.Value,
		HrAllowed:              plan.HrAllowed.Value,
		RaAllowed:              plan.RaAllowed.Value,
		NwkGeoLoc:              plan.NwkGeoLoc.Value,
		TargetPer:              uint32(plan.TargetPer.Value),
		MinGwDiversity:         uint32(plan.MinGwDiversity.Value),
		GwsPrivate:             plan.GwsPrivate.Value,
	}
	request := api.CreateDeviceProfileRequest{
		DeviceProfile: &deviceprofile,
	}

	client := api.NewDeviceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Create(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating deviceprofile",
			err.Error(),
		)
		return
	}

	plan.ID = types.String{Value: response.Id}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	LoadRespFromResourceRead(ctx, NewCreateResponse(resp), r, req.ProviderMeta)
}

// Read resource information
func (r resourceDeviceProfile) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Get current state
	var state DeviceProfile
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.GetDeviceProfileRequest{
		Id: state.ID.Value,
	}

	client := api.NewDeviceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Get(ctx, &request)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError(
			"Error reading device profile",
			err.Error(),
		)
		return
	}

	state.Name = types.String{Value: response.DeviceProfile.Name}
	state.OrganizationId = types.Int64{Value: int64(response.DeviceProfile.OrganizationId)}
	state.NetworkServerId = types.Int64{Value: int64(response.DeviceProfile.NetworkServerId)}
	state.UlRate = types.Int64{Value: int64(response.ServiceProfile.UlRate)}
	state.UlBucketSize = types.Int64{Value: int64(response.ServiceProfile.UlBucketSize)}
	state.UlRatePolicy = types.Int64{Value: int64(response.ServiceProfile.UlRatePolicy)}
	state.DlRate = types.Int64{Value: int64(response.ServiceProfile.DlRate)}
	state.DlBucketSize = types.Int64{Value: int64(response.ServiceProfile.DlBucketSize)}
	state.DlRatePolicy = types.Int64{Value: int64(response.ServiceProfile.DlRatePolicy)}
	state.AddGwMetadata = types.Bool{Value: response.ServiceProfile.AddGwMetadata}
	state.DevStatusReqFreq = types.Int64{Value: int64(response.ServiceProfile.DevStatusReqFreq)}
	state.ReportDevStatusBattery = types.Bool{Value: response.ServiceProfile.ReportDevStatusBattery}
	state.ReportDevStatusMargin = types.Bool{Value: response.ServiceProfile.ReportDevStatusMargin}
	state.DrMin = types.Int64{Value: int64(response.ServiceProfile.DrMin)}
	state.DrMax = types.Int64{Value: int64(response.ServiceProfile.DrMax)}
	state.ChannelMask = types.Int64{Null: true} // TODO
	state.PrAllowed = types.Bool{Value: response.ServiceProfile.PrAllowed}
	state.HrAllowed = types.Bool{Value: response.ServiceProfile.HrAllowed}
	state.RaAllowed = types.Bool{Value: response.ServiceProfile.RaAllowed}
	state.NwkGeoLoc = types.Bool{Value: response.ServiceProfile.NwkGeoLoc}
	state.TargetPer = types.Int64{Value: int64(response.ServiceProfile.TargetPer)}
	state.MinGwDiversity = types.Int64{Value: int64(response.ServiceProfile.MinGwDiversity)}
	state.GwsPrivate = types.Bool{Value: response.ServiceProfile.GwsPrivate}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceDeviceProfile) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve values from plan
	var plan DeviceProfile
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state DeviceProfile
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = state.ID

	deviceprofile := api.DeviceProfile{
		Id:                     plan.ID.Value,
		Name:                   plan.Name.Value,
		OrganizationId:         plan.OrganizationId.Value,
		NetworkServerId:        plan.NetworkServerId.Value,
		UlRate:                 uint32(plan.UlRate.Value),
		UlBucketSize:           uint32(plan.UlBucketSize.Value),
		UlRatePolicy:           api.RatePolicy(plan.UlRatePolicy.Value),
		DlRate:                 uint32(plan.DlRate.Value),
		DlBucketSize:           uint32(plan.DlBucketSize.Value),
		DlRatePolicy:           api.RatePolicy(plan.DlRatePolicy.Value),
		AddGwMetadata:          plan.AddGwMetadata.Value,
		DevStatusReqFreq:       uint32(plan.DevStatusReqFreq.Value),
		ReportDevStatusBattery: plan.ReportDevStatusBattery.Value,
		ReportDevStatusMargin:  plan.ReportDevStatusMargin.Value,
		DrMin:                  uint32(plan.DrMin.Value),
		DrMax:                  uint32(plan.DrMax.Value),
		ChannelMask:            nil, // TODO
		PrAllowed:              plan.PrAllowed.Value,
		HrAllowed:              plan.HrAllowed.Value,
		RaAllowed:              plan.RaAllowed.Value,
		NwkGeoLoc:              plan.NwkGeoLoc.Value,
		TargetPer:              uint32(plan.TargetPer.Value),
		MinGwDiversity:         uint32(plan.MinGwDiversity.Value),
		GwsPrivate:             plan.GwsPrivate.Value,
	}
	request := api.UpdateDeviceProfileRequest{
		DeviceProfile: &deviceprofile,
	}

	client := api.NewDeviceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Update(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating device profile",
			err.Error(),
		)
		return
	}

	LoadRespFromResourceRead(ctx, NewUpdateResponse(resp), r, req.ProviderMeta)
}

// Delete resource
func (r resourceDeviceProfile) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Get current state
	var state DeviceProfile
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.DeleteDeviceProfileRequest{
		Id: state.ID.Value,
	}

	client := api.NewDeviceProfileServiceClient(r.p.Conn(ctx))
	resp.Diagnostics.Append(r.p.Diagnostics()...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := client.Delete(ctx, &request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device profile",
			err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r resourceDeviceProfile) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req.ID)

	LoadRespFromResourceRead(ctx, NewImportResponse(resp), r, tfsdk.Config{})
}
*/
