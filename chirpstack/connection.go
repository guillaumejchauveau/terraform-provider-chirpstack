package chirpstack

import (
	"context"
	"fmt"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc"
)

type APIToken string

func (a APIToken) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", a),
	}, nil
}

func (a APIToken) RequireTransportSecurity() bool {
	return false
}

type ConnectionData struct {
	Server   types.String `tfsdk:"server"`
	Key      types.String `tfsdk:"key"`
	Email    types.String `tfsdk:"email"`
	Password types.String `tfsdk:"password"`
}

func (p *provider) dial() (*grpc.ClientConn, error) {
	if p.ConnectionData.Server.Unknown ||
		p.ConnectionData.Key.Unknown ||
		p.ConnectionData.Email.Unknown ||
		p.ConnectionData.Password.Unknown {
		return nil, fmt.Errorf("configuration error: connection has unknown properties")
	}

	if p.ConnectionData.Key.Null {
		if p.ConnectionData.Email.Null || p.ConnectionData.Password.Null {
			return nil, fmt.Errorf("configuration error: either key or email/password must be set")
		}

		dialOpts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithInsecure(), // remove this when using TLS
		}
		conn, err := grpc.Dial(p.ConnectionData.Server.Value, dialOpts...)
		if err != nil {
			return nil, err
		}

		loginReq := api.LoginRequest{
			Email:    p.ConnectionData.Email.Value,
			Password: p.ConnectionData.Password.Value,
		}

		internalClient := api.NewInternalServiceClient(conn)
		loginResp, err := internalClient.Login(p.ctx, &loginReq)
		conn.Close()
		if err != nil {
			return nil, err
		}
		p.ConnectionData.Key = types.String{Value: loginResp.Jwt}
	}

	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(APIToken(p.ConnectionData.Key.Value)),
		grpc.WithInsecure(), // remove this when using TLS
	}

	return grpc.Dial(p.ConnectionData.Server.Value, dialOpts...)
}

func (p *provider) Conn() *grpc.ClientConn {
	if p.conn == nil {
		conn, err := p.dial()
		if err != nil {
			p.Diagnostics.AddError(
				"Error establishing connection",
				"Could not connect to API: "+err.Error(),
			)
			return nil
		}
		p.conn = conn
	}
	return p.conn
}
