package models

import (
	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func UserSchema() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id":          {Type: types.Int64Type, Computed: true},
			"email":       {Type: types.StringType, Required: true},
			"password":    {Type: types.StringType, Required: true, Sensitive: true},
			"is_active":   {Type: types.BoolType, Optional: true, Computed: true},
			"is_admin":    {Type: types.BoolType, Optional: true, Computed: true},
			"note":        {Type: types.StringType, Optional: true, Computed: true},
			"session_ttl": {Type: types.Int64Type, Optional: true, Computed: true},
		},
	}
}

type User struct {
	Id         types.Int64  `tfsdk:"id"`
	Email      types.String `tfsdk:"email"`
	Password   types.String `tfsdk:"password"`
	IsActive   types.Bool   `tfsdk:"is_active"`
	IsAdmin    types.Bool   `tfsdk:"is_admin"`
	Note       types.String `tfsdk:"note"`
	SessionTtl types.Int64  `tfsdk:"session_ttl"`
}

func UserFromApiType(s *api.User, password string) User {
	return User{
		Id:         types.Int64{Value: int64(s.Id)},
		Email:      types.String{Value: s.Email},
		Password:   types.String{Value: password},
		IsActive:   types.Bool{Value: s.IsActive},
		IsAdmin:    types.Bool{Value: s.IsAdmin},
		Note:       types.String{Value: s.Note},
		SessionTtl: types.Int64{Value: int64(s.SessionTtl)},
	}
}

func (s *User) ToApiType() api.User {
	return api.User{
		Id:         s.Id.Value,
		Email:      s.Email.Value,
		IsActive:   s.IsActive.Value,
		IsAdmin:    s.IsAdmin.Value,
		Note:       s.Note.Value,
		SessionTtl: int32(s.SessionTtl.Value),
	}
}

func (u *User) Equal(o User) bool {
	return u.Email.Value == o.Email.Value &&
		u.Password.Value == o.Password.Value &&
		u.IsActive.Value == o.IsActive.Value &&
		u.IsAdmin.Value == o.IsAdmin.Value &&
		u.Note.Value == o.Note.Value &&
		u.SessionTtl.Value == o.SessionTtl.Value
}
