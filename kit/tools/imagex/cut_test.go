package imagex

import (
	"testing"
)

func TestCutImage(t *testing.T) {
	type args struct {
		inputPath string
		outputDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"testdata/cat_00000_00.png", "testdata"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CutImage(tt.args.inputPath, tt.args.outputDir); (err != nil) != tt.wantErr {
				t.Errorf("cutImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
