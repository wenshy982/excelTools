package imagex

import (
	"testing"
)

func TestConvertWebpToPng(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"testdata/convert.webp"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertWebpToPng(tt.args.path, "testdata/convert.png"); (err != nil) != tt.wantErr {
				t.Errorf("convertToPng() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
