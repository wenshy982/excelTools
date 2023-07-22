package imagex

import (
	"testing"
)

func TestMergeImages(t *testing.T) {
	type args struct {
		img1 string
		img2 string
		img3 string
		img4 string
		dir  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test1",
			args{
				img1: "testdata/1.png",
				img2: "testdata/2.png",
				img3: "testdata/3.png",
				img4: "testdata/4.png",
				dir:  "testdata",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = MergeImages(tt.args.img1, tt.args.img2, tt.args.img3, tt.args.img4, tt.args.dir)
		})
	}
}
