package day5

import (
	"reflect"
	"testing"
)

const testInput = `0
3
0
1
-3`

func Test_loadInput(t *testing.T) {
	type args struct {
		steps string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test",
			args{testInput},
			[]int{0, 3, 0, 1, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadInput(tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_howManySteps(t *testing.T) {
	type args struct {
		steps []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test input",
			args{[]int{0, 3, 0, 1, -3}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManySteps(tt.args.steps); got != tt.want {
				t.Errorf("howManySteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testStepsInput(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"Input.txt",
			396086,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testStepsInput(); got != tt.want {
				t.Errorf("testStepsInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_howManyStepsWithCaveats(t *testing.T) {
	type args struct {
		steps []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"same input",
			args{[]int{0, 3, 0, 1, -3}},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManyStepsWithCaveats(tt.args.steps); got != tt.want {
				t.Errorf("howManyStepsWithCaveats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testStepsInputPart2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"test part 2",
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testStepsInputPart2(); got != tt.want {
				t.Errorf("testStepsInputPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
