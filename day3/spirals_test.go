package day3

import (
	"reflect"
	"testing"
)

func TestGenerateSpiral(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSpiral(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSpiral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStepsFromSpiral(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{1},
			0,
		},
		{
			"12",
			args{12},
			3,
		},
		{
			"23",
			args{23},
			2,
		},
		{
			"1024",
			args{1024},
			31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StepsFromSpiral(tt.args.num); got != tt.want {
				t.Errorf("StepsFromSpiral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionLocation(t *testing.T) {
	type args struct {
		pos int
	}
	tests := []struct {
		name string
		args args
		want position
	}{
		{
			"1",
			args{1},
			position{0, 0},
		},
		{
			"12",
			args{12},
			position{1, 1},
		},
		{
			"23",
			args{23},
			position{0, -2},
		},
		{
			"1024",
			args{1024},
			position{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PositionLocation(tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PositionLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
