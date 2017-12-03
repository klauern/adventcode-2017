package day2

import (
	"reflect"
	"testing"
)

func TestDifferenceLargestSmallest(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"5 1 9 5",
			args{"5	1	9	5"},
			8,
		},
		{
			"7 5 3",
			args{"7	5	3"},
			4,
		},
		{
			"2 4 6 8",
			args{"2	4	6	8"},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceLargestSmallest(tt.args.line); got != tt.want {
				t.Errorf("DifferenceLargestSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLine(t *testing.T) {
	type args struct {
		line    string
		splitBy string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"7 5 3",
			args{"7 5 3", " "},
			[]int{7, 5, 3},
		},
		{
			"2 4 6 8",
			args{"2 4 6 8", " "},
			[]int{2, 4, 6, 8},
		},
		{
			"5 1 9 5",
			args{"5 1 9 5", " "},
			[]int{5, 1, 9, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLine(tt.args.line, tt.args.splitBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksumTestSpreadsheet(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"sample test",
			18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChecksumTestSpreadsheet(); got != tt.want {
				t.Errorf("ChecksumTestSpreadsheet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksumInput(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"sample",
			41887,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChecksumInput(); got != tt.want {
				t.Logf("ChecksumInput() = %v", got)
				t.Errorf("ChecksumInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checksumString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"spreadsheet",
			args{spreadsheet},
			18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksumString(tt.args.str); got != tt.want {
				t.Errorf("checksumString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findEvenlyDivisible(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"5 9 2 8",
			args{"5	9	2	8"},
			[]int{2, 8},
		},
		{
			"9 4 7 3",
			args{"9	4	7	3"},
			[]int{9, 3},
		},
		{
			"3 8 6 5",
			args{"3	8	6	5"},
			[]int{3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEvenlyDivisible(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEvenlyDivisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDivisibleChecksumForLine(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"5 9 2 8",
			args{"5	9	2	8"},
			4,
		},
		{
			"9 4 7 3",
			args{"9	4	7	3"},
			3,
		},
		{
			"3 8 6 5",
			args{"3	8	6	5"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDivisibleChecksumForLine(tt.args.str); got != tt.want {
				t.Errorf("getDivisibleChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDivisibleChecksumFromInput(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"input",
			226,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDivisibleChecksumFromInput(); got != tt.want {
				t.Errorf("getDivisibleChecksumFromInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
