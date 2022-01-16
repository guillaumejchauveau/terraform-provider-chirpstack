package models

import (
	"context"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ApplicationHttpSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"application_id": {Type: types.Int64Type, Required: true},
			"headers": {
				Type:     types.MapType{ElemType: types.StringType},
				Optional: true,
				Computed: true,
			},
			"marshaler":          {Type: types.StringType, Required: true},
			"event_endpoint_url": {Type: types.StringType, Required: true},
		},
	}
}

type ApplicationHttp struct {
	ApplicationId    types.Int64  `tfsdk:"application_id"`
	Headers          types.Map    `tfsdk:"headers"`
	Marshaler        types.String `tfsdk:"marshaler"`
	EventEndpointUrl types.String `tfsdk:"event_endpoint_url"`
}

func ApplicationHttpFromApiType(s *api.HTTPIntegration) ApplicationHttp {
	headers := map[string]attr.Value{}
	for _, v := range s.Headers {
		headers[v.Key] = types.String{Value: v.Value}
	}
	return ApplicationHttp{
		ApplicationId:    types.Int64{Value: s.ApplicationId},
		Headers:          types.Map{Elems: headers, ElemType: types.StringType},
		Marshaler:        types.String{Value: s.Marshaler.String()},
		EventEndpointUrl: types.String{Value: s.EventEndpointUrl},
	}
}

func (s *ApplicationHttp) ToApiType(ctx context.Context) api.HTTPIntegration {
	headers := []*api.HTTPIntegrationHeader{}
	for k, v := range s.Headers.Elems {
		value, err := v.ToTerraformValue(ctx)
		if err != nil {
			panic(err)
		}
		headers = append(headers, &api.HTTPIntegrationHeader{Key: k, Value: value.(string)})
	}
	return api.HTTPIntegration{
		ApplicationId:    s.ApplicationId.Value,
		Headers:          headers,
		Marshaler:        api.Marshaler(api.Marshaler_value[s.Marshaler.Value]),
		EventEndpointUrl: s.EventEndpointUrl.Value,
	}
}
