package xls

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	type args struct {
		s         []string
		v         string
		skipIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s:         []string{"SG", "MY", "TH", "VN", "PH", "TW", "BR", "MX", "CO", "CL"},
				v:         "SG",
				skipIndex: 0,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				s:         []string{"SIP卖家结算货币", "结算价货币", "店铺类型", "Region", "SIP产品\n活动结算价"},
				v:         "SIP产品\n活动结算价",
				skipIndex: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.s, tt.args.v, tt.args.skipIndex); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
