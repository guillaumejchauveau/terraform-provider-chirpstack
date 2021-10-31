package models

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"google.golang.org/protobuf/types/known/durationpb"
)

var _ attr.Type = DurationType{}

type DurationType struct {
}

func (t DurationType) TerraformType(context.Context) tftypes.Type {
	return tftypes.String
}

func (t DurationType) ValueFromTerraform(_ context.Context, in tftypes.Value) (attr.Value, error) {
	if !in.IsKnown() {
		return Duration{Unknown: true}, nil
	}
	if in.IsNull() {
		return Duration{Null: true}, nil
	}
	var s string
	err := in.As(&s)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`(\d+h)?(\d+m)?(\d+s)?`)
	match := r.FindStringSubmatch(s)
	s_h := strings.TrimSuffix(match[1], "h")
	s_m := strings.TrimSuffix(match[2], "m")
	s_s := strings.TrimSuffix(match[3], "s")
	if s_h == "" {
		s_h = "0"
	}
	if s_m == "" {
		s_m = "0"
	}
	if s_s == "" {
		s_s = "0"
	}
	hours, err := strconv.ParseUint(s_h, 10, 64)
	if err != nil {
		return nil, err
	}
	minutes, err := strconv.ParseUint(s_m, 10, 64)
	if err != nil {
		return nil, err
	}
	seconds, err := strconv.ParseUint(s_s, 10, 64)
	if err != nil {
		return nil, err
	}
	return Duration{
		Value: durationpb.New(
			time.Duration(hours)*time.Hour +
				time.Duration(minutes)*time.Minute +
				time.Duration(seconds)*time.Second),
	}, nil
}

func (t DurationType) Equal(o attr.Type) bool {
	return t == o
}

func (t DurationType) String() string {
	return "DurationType"
}

func (t DurationType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return nil, fmt.Errorf("cannot apply AttributePathStep %T to %s", step, t.String())
}

var _ attr.Value = Duration{}

type Duration struct {
	// Unknown will be true if the value is not yet known.
	Unknown bool

	// Null will be true if the value was not set, or was explicitly set to
	// null.
	Null bool

	// Value contains the set value, as long as Unknown and Null are both
	// false.
	Value *duration.Duration
}

func (d Duration) Type(context.Context) attr.Type {
	return DurationType{}
}

func (d Duration) ToTerraformValue(context.Context) (interface{}, error) {
	return d.Value.AsDuration().Truncate(time.Second).String(), nil
}

func (d Duration) Equal(other attr.Value) bool {
	o, ok := other.(Duration)
	if !ok {
		return false
	}
	if d.Unknown != o.Unknown {
		return false
	}
	if d.Null != o.Null {
		return false
	}
	return d.Value.AsDuration().Truncate(time.Second) == o.Value.AsDuration().Truncate(time.Second)
}
