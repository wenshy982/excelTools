package qrcodex

import (
	"os"
	"testing"
)

func TestBase64CodeUrl(t *testing.T) {
	type args struct {
		codeUrl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				codeUrl: "https://www.baidu.com",
			},
			want: "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABVUlEQVR42uyYwa0zMQiEZ+WDjy6BUtxaSnMpLsFHHyLzC/g32s1LUgB4Dk9Pm++EzMCAra2trV9i1TPxzBNllE725eELWPInPdPMMw+gdOrypYYDDimTAlKnAZJStaiA/l8GEBxI0ixDG8cnYH2BPJF5fG8c78Dpk+oPv4zUOfBSlheD0r/OR9/AwmGV4QkUZu6gVm9z0wOAgxeQ5GetBNmDaAgGHLYksaxJZRTuxL22iz+EAZa1xYT5A6jXqz8EAZgXEvPT7OHjvPAArJcNXvwBtSEYcElJuikRm0E8ggHr4KVLcRarhBhEr2/7ZBRA06LGAytbq5d54QP4L3sPtgVRROB2JDGfpPYpHTgHziMJ61gUm1Sf5HDAeSSxddEaxy9g6WAAuh6EBXRBGEXSAf6MxRDAeSQ5CwXitzgZAngdSXRsSkwCbinJB7C1tRVR/wIAAP//mXiC2Cv0frEAAAAASUVORK5CYII=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64CodeUrl(tt.args.codeUrl); got != tt.want {
				t.Errorf("Base64CodeUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQRCode(t *testing.T) {
	r, err := QRCode("https://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestQRCodeToFile(t *testing.T) {
	filename := "test.png"
	err := QRCodeToFile("https://www.baidu.com", filename)
	if err != nil {
		t.Error(err)
	}
	_ = os.Remove(filename)
}
