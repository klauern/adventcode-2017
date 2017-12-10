package day6

import (
	"reflect"
	"testing"
)

func TestNumReallocsCycle(t *testing.T) {
	type args struct {
		mem memory
	}
	tests := []struct {
		name string
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumReallocsCycle(tt.args.mem); got != tt.want {
				t.Errorf("NumReallocsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reallocateSegments(t *testing.T) {
	type args struct {
		mem memory
	}
	tests := []struct {
		name string
		args args
		want memory
	}{
		{
			"step 1",
			args{memory{0, 2, 7, 0}},
			memory{2, 4, 1, 2},
		},
		{
			"step 2",
			args{memory{2, 4, 1, 2}},
			memory{3, 1, 2, 3},
		},
		{
			"step 3",
			args{memory{3, 1, 2, 3}},
			memory{0, 2, 3, 4},
		},
		{
			"step 4",
			args{memory{0, 2, 3, 4}},
			memory{1, 3, 4, 1},
		},
		{
			"step 5 (first repetition)",
			args{memory{1, 3, 4, 1}},
			memory{2, 4, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reallocateSegments(tt.args.mem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reallocateSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memory_maxIdx(t *testing.T) {
	tests := []struct {
		name string
		m    memory
		want int
	}{
		{
			"step 1",
			memory{0, 2, 3, 4},
			3,
		},
		{
			"step 2",
			memory{2, 4, 1, 2},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.maxIdx(); got != tt.want {
				t.Errorf("memory.maxIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}
