package day6

import "testing"

func TestNumReallocsCycle(t *testing.T) {
	type args struct {
		mem []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumReallocsCycle(tt.args.mem); got != tt.want {
				t.Errorf("NumReallocsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
