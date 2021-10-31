package chirpstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type Response interface {
	State() tfsdk.State
	SetState(tfsdk.State)
	Diagnostics() diag.Diagnostics
	SetDiagnostics(diag.Diagnostics)
}

func LoadRespFromResourceRead(ctx context.Context, resp Response, r tfsdk.Resource, ProviderMeta tfsdk.Config) {
	readResp := tfsdk.ReadResourceResponse{
		State:       resp.State(),
		Diagnostics: resp.Diagnostics(),
	}
	r.Read(ctx, tfsdk.ReadResourceRequest{
		State:        resp.State(),
		ProviderMeta: ProviderMeta,
	}, &readResp)
	resp.SetState(readResp.State)
	resp.SetDiagnostics(readResp.Diagnostics)
}

type CreateResponse struct {
	resp *tfsdk.CreateResourceResponse
}

func NewCreateResponse(resp *tfsdk.CreateResourceResponse) CreateResponse {
	return CreateResponse{
		resp: resp,
	}
}

func (resp CreateResponse) State() tfsdk.State {
	return resp.resp.State
}
func (resp CreateResponse) SetState(state tfsdk.State) {
	resp.resp.State = state
}
func (resp CreateResponse) Diagnostics() diag.Diagnostics {
	return resp.resp.Diagnostics
}
func (resp CreateResponse) SetDiagnostics(diags diag.Diagnostics) {
	resp.resp.Diagnostics = diags
}

type UpdateResponse struct {
	resp *tfsdk.UpdateResourceResponse
}

func NewUpdateResponse(resp *tfsdk.UpdateResourceResponse) UpdateResponse {
	return UpdateResponse{
		resp: resp,
	}
}

func (resp UpdateResponse) State() tfsdk.State {
	return resp.resp.State
}
func (resp UpdateResponse) SetState(state tfsdk.State) {
	resp.resp.State = state
}
func (resp UpdateResponse) Diagnostics() diag.Diagnostics {
	return resp.resp.Diagnostics
}
func (resp UpdateResponse) SetDiagnostics(diags diag.Diagnostics) {
	resp.resp.Diagnostics = diags
}

type ImportResponse struct {
	resp *tfsdk.ImportResourceStateResponse
}

func NewImportResponse(resp *tfsdk.ImportResourceStateResponse) ImportResponse {
	return ImportResponse{
		resp: resp,
	}
}

func (resp ImportResponse) State() tfsdk.State {
	return resp.resp.State
}
func (resp ImportResponse) SetState(state tfsdk.State) {
	resp.resp.State = state
}
func (resp ImportResponse) Diagnostics() diag.Diagnostics {
	return resp.resp.Diagnostics
}
func (resp ImportResponse) SetDiagnostics(diags diag.Diagnostics) {
	resp.resp.Diagnostics = diags
}
