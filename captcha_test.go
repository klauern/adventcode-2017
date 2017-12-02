package day1

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1122",
			args: args{"1122"},
			want: 3,
		},
		{
			"1111",
			args{"1111"},
			4,
		},
		{
			"1234",
			args{"1234"},
			0,
		},
		{
			"91212129",
			args{"91212129"},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.seq); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
