package xls

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func interfaceEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a != nil && b != nil {
		ai, aIsInt := a.(int)
		bi, bIsInt := b.(int)
		if aIsInt && bIsInt && ai == bi {
			return true
		}

		af, aIsFloat := a.(float64)
		bf, bIsFloat := b.(float64)
		if aIsFloat && bIsFloat && math.Abs(af-bf) <= epsilon {
			return true
		}

		as, aIsString := a.(string)
		bs, bIsString := b.(string)
		if aIsString && bIsString && as == bs {
			return true
		}
	}

	return false
}

func Equal(got []interface{}, want []interface{}) bool {
	if len(got) != len(want) {
		return false
	}
	for i, v := range got {
		if !interfaceEqual(v, want[i]) {
			return false
		}
	}
	return true
}

func TestExcel_NewRow(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "int test",
			args: args{
				values: []string{"1", "2", "3"},
			},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "test2",
			args: args{
				values: []string{"1", "2", "3.1"},
			},
			want: []interface{}{1, 2, 3.1},
		},
		{
			name: "float test",
			args: args{
				values: []string{"1.1", "2.2", "3.3"},
			},
			want: []interface{}{1.1, 2.2, 3.3},
		},
		{
			name: "empty test",
			args: args{
				values: []string{"", "", ""},
			},
			want: []interface{}{nil, nil, nil},
		},
		{
			name: "string test",
			args: args{
				values: []string{"test1", "test2", "test3"},
			},
			want: []interface{}{"test1", "test2", "test3"},
		},
		{
			name: "all test",
			args: args{
				values: []string{
					"1", "2", "3",
					"4.1", "5.2", "6.3",
					"", "", "",
					"test1", "test2", "test3", "test4", "test5", "test6",
				},
			},
			want: []interface{}{
				1, 2, 3,
				4.1, 5.2, 6.3,
				nil, nil, nil,
				"test1", "test2", "test3", "test4", "test5", "test6",
			},
		},
	}
	var e = New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := e.NewRow(tt.args.values...); !Equal(got, tt.want) {
				t.Errorf("NewRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkExcel_NewRow(b *testing.B) {
	var e = New()
	for i := 0; i < b.N; i++ {
		e.NewRow(
			"1", "2", "3",
			"4.1", "5.2", "6.3",
			"", "", "",
			"test1", "test2", "test3", "test4", "test5", "test6",
		)
	}
}

func TestExcel_IntCol(t *testing.T) {
	var e = New(WithErrors("testdata/errors.xlsx"))
	f := e.OpenErrors()
	defer func() { _ = f.Close() }()
	col := e.IntCol(f, 1, 1)
	for _, v := range col {
		if v == "" {
			t.Error("IntCol() error, got empty string")
		}
		if !isInt(v) {
			t.Errorf("IntCol() error, got %v", v)
		}
	}
}

func BenchmarkExcel_IntCol(b *testing.B) {
	var e = New(WithErrors("testdata/errors.xlsx"))
	f := e.OpenErrors()
	defer func() { _ = f.Close() }()
	for i := 0; i < b.N; i++ {
		e.IntCol(f, 1, 1)
	}
}

func TestExcelColumnIndex(t *testing.T) {
	type args struct {
		values string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{
				values: "A",
			},
			want: 0,
		},
		{
			name: "Z",
			args: args{
				values: "Z",
			},
			want: 25,
		},
		{
			name: "AA",
			args: args{
				values: "AA",
			},
			want: 26,
		},
		{
			name: "AZ",
			args: args{
				values: "AZ",
			},
			want: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colIndex(tt.args.values); got != tt.want {
				t.Errorf("ExcelColumnIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkExcelColumnIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colIndex("AZ")
	}
}
