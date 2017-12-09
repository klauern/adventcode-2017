package day6

import (
	"reflect"
	"testing"
)

func TestNumReallocsCycle(t *testing.T) {
	type args struct {
		mem []int
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
		mem []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"step 1",
			args{[]int{0, 2, 7, 0}},
			[]int{2, 4, 1, 2},
		},
		{
			"step 2",
			args{[]int{2, 4, 1, 2}},
			[]int{3, 1, 2, 3},
		},
		{
			"step 3",
			args{[]int{3, 1, 2, 3}},
			[]int{0, 2, 3, 4},
		},
		{
			"step 4",
			args{[]int{0, 2, 3, 4}},
			[]int{1, 3, 4, 1},
		},
		{
			"step 5 (first repetition)",
			args{[]int{1, 3, 4, 1}},
			[]int{2, 4, 1, 2},
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
