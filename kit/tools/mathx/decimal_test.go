package mathx

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestFloatAdd(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{1.23000000001, 1.23000000001}, 2.46},
		{"test2", args{1.23111111111, 1.23111111111}, 2.46},
		{"test3", args{1.23222222222, 1.23222222222}, 2.46},
		{"test4", args{1.23333333333, 1.23333333333}, 2.47},
		{"test5", args{1.23444444444, 1.23444444444}, 2.47},
		{"test6", args{1.23555555555, 1.23555555555}, 2.47},
		{"test7", args{1.23666666666, 1.23666666666}, 2.47},
		{"test8", args{1.23777777777, 1.23777777777}, 2.48},
		{"test9", args{1.23888888888, 1.23888888888}, 2.48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("FloatAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFloatAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatAdd(1.23000000001, 1.23000000001)
		FloatAdd(1.23888888888, 1.23888888888)
	}
}

func TestFloatSub(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{1.23999999999, 1.23000000001}, 0.01},
		{"test2", args{1.23888888888, 1.23111111111}, 0.01},
		{"test3", args{1.23777777777, 1.23222222222}, 0.01},
		{"test4", args{1.23666666666, 1.23333333333}, 0},
		{"test5", args{1.23555555555, 1.23444444444}, 0},
		{"test6", args{1.23444444444, 1.23555555555}, 0},
		{"test7", args{1.23333333333, 1.23666666666}, 0},
		{"test8", args{1.23222222222, 1.23777777777}, -0.01},
		{"test9", args{1.23111111111, 1.23888888888}, -0.01},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatSub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("FloatAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFloatSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatSub(1.23000000001, 1.23000000001)
		FloatSub(1.23888888888, 1.23888888888)
	}
}

func TestFloatMul(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{1.23000000001, 1.23000000001}, 1.51},
		{"test2", args{1.23888888888, 1.23888888888}, 1.53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatMul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("FloatMul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFloatMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatMul(1.23000000001, 1.23000000001)
		FloatMul(1.23888888888, 1.23888888888)
	}
}

func TestFloatDiv(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{1.23000000001, 1.23000000001}, 1},
		{"test2", args{1.23888888888, 1.23888888888}, 1},
		{"test3", args{1.23888888888, 1.23000000001}, 1.01},
		{"test4", args{1.23000000001, 1.23888888888}, 0.99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatDiv(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("FloatDiv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFloatDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatDiv(1.23000000001, 1.23000000001)
		FloatDiv(1.23888888888, 1.23888888888)
	}
}

func TestFloatRound(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{1.23000000001}, 1.23},
		{"test2", args{1.23111111111}, 1.23},
		{"test3", args{1.23222222222}, 1.23},
		{"test4", args{1.23333333333}, 1.23},
		{"test5", args{1.23444444444}, 1.23},
		{"test6", args{1.23555555555}, 1.24},
		{"test7", args{1.23666666666}, 1.24},
		{"test8", args{1.23777777777}, 1.24},
		{"test9", args{1.23888888888}, 1.24},
		{"test10", args{1.23999999999}, 1.24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatRound(tt.args.value); got != tt.want {
				t.Errorf("FloatRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFloatRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatRound(1.23000000001)
		FloatRound(1.23999999999)
	}
}

func TestDecimalRound(t *testing.T) {
	type args struct {
		value decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{decimal.NewFromFloat(1.23000000001)}, 1.23},
		{"test2", args{decimal.NewFromFloat(1.23111111111)}, 1.23},
		{"test3", args{decimal.NewFromFloat(1.23222222222)}, 1.23},
		{"test4", args{decimal.NewFromFloat(1.23333333333)}, 1.23},
		{"test5", args{decimal.NewFromFloat(1.23444444444)}, 1.23},
		{"test6", args{decimal.NewFromFloat(1.23555555555)}, 1.24},
		{"test7", args{decimal.NewFromFloat(1.23666666666)}, 1.24},
		{"test8", args{decimal.NewFromFloat(1.23777777777)}, 1.24},
		{"test9", args{decimal.NewFromFloat(1.23888888888)}, 1.24},
		{"test10", args{decimal.NewFromFloat(1.23999999999)}, 1.24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalRound(tt.args.value); got != tt.want {
				t.Errorf("DecimalRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDecimalRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecimalRound(decimal.NewFromFloat(1.23000000001))
		DecimalRound(decimal.NewFromFloat(1.23999999999))
	}
}

func TestFenToPoints(t *testing.T) {
	type args struct {
		money int64
		ratio int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				money: 100,
				ratio: 10,
			},
			want: 10,
		},
		{
			name: "test2",
			args: args{
				money: 10,
				ratio: 10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FenToPoints(tt.args.money, tt.args.ratio); got != tt.want {
				t.Errorf("FenToPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFenToPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FenToPoints(100, 10)
	}
}

func TestYuanToPoints(t *testing.T) {
	type args struct {
		money float64
		ratio int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				money: 100,
				ratio: 10,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YuanToPoints(tt.args.money, tt.args.ratio); got != tt.want {
				t.Errorf("YuanToPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkYuanToPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YuanToPoints(100, 10)
	}
}

func TestPointsToFen(t *testing.T) {
	type args struct {
		points float64
		ratio  int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				points: 2.84,
				ratio:  10,
			},
			want: 28,
		},
		{
			name: "test2",
			args: args{
				points: 50.22,
				ratio:  10,
			},
			want: 502,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointsToFen(tt.args.points, tt.args.ratio); got != tt.want {
				t.Errorf("PointsToFen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPointsToFen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PointsToFen(10, 10)
	}

}

func TestPointsToYuan(t *testing.T) {
	type args struct {
		points float64
		ratio  int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				points: 2.84,
				ratio:  10,
			},
			want: 0.28,
		},
		{
			name: "test2",
			args: args{
				points: 50.55,
				ratio:  5,
			},
			want: 10.11,
		},
		{
			name: "test3",
			args: args{
				points: 2100,
				ratio:  10,
			},
			want: 210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointsToYuan(tt.args.points, tt.args.ratio); got != tt.want {
				t.Errorf("PointsToYuan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPointsToYuan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PointsToYuan(10, 10)
	}
}

func TestCharacterToPoints(t *testing.T) {
	type args struct {
		characterCount int32
		ratio          int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				characterCount: 2,
				ratio:          500,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				characterCount: 3,
				ratio:          500,
			},
			want: 0.01,
		},
		{
			name: "test3",
			args: args{
				characterCount: 32,
				ratio:          500,
			},
			want: 0.06,
		},
		{
			name: "test4",
			args: args{
				characterCount: 100,
				ratio:          500,
			},
			want: 0.2,
		},
		{
			name: "test5",
			args: args{
				characterCount: 1700,
				ratio:          500,
			},
			want: 3.4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharacterToPoints(tt.args.characterCount, tt.args.ratio); got != tt.want {
				t.Errorf("CharacterToPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCharacterToPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CharacterToPoints(100, 500)
	}
}
