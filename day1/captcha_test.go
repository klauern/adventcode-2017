package day1

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSumSequence(t *testing.T) {
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
			if got := SumSequence(tt.args.seq); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumChars(t *testing.T) {
	type args struct {
		prev rune
		cur  rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"11",
			args{rune("11"[0]), rune("11"[1])},
			2,
		},
		{
			"23",
			args{rune("23"[0]), rune("23"[1])},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumChars(tt.args.prev, tt.args.cur); got != tt.want {
				t.Errorf("SumChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSequenceWithInput(t *testing.T) {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sum is %d", SumSequence(string(b)))
	fmt.Printf("Sum of 'input' is %d\n", SumSequence(string(b)))
}

func TestSumRange(t *testing.T) {
	type args struct {
		summable []int
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
			if got := SumRange(tt.args.summable); got != tt.want {
				t.Errorf("SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSame(t *testing.T) {
	type args struct {
		prev rune
		cur  rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"ab no",
			args{'a', 'b'},
			false,
		},
		{
			"aa yes",
			args{'a', 'a'},
			true,
		},
		{
			"12 no",
			args{'1', '2'},
			false,
		},
		{
			"11 yes",
			args{'1', '1'},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSame(tt.args.prev, tt.args.cur); got != tt.want {
				t.Errorf("IsSame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		digit rune
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
			if got := Convert(tt.args.digit); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOverRange(t *testing.T) {
	type args struct {
		seq  string
		jump int
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
			if got := SumOverRange(tt.args.seq, tt.args.jump); got != tt.want {
				t.Errorf("SumOverRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSequenceMod2(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1212",
			args{"1212"},
			6,
		},
		{
			"1221",
			args{"1221"},
			0,
		},
		{
			"123425",
			args{"123425"},
			4,
		},
		{
			"123123",
			args{"123123"},
			12,
		},
		{
			"12131415",
			args{"12131415"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSequenceMod2(tt.args.seq); got != tt.want {
				t.Errorf("SumSequenceMod2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSequenceMod2WithInput(t *testing.T) {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sum is %d", SumSequenceMod2(string(b)))
	fmt.Printf("Sum of 'input' for Mod2 is %d\n", SumSequenceMod2(string(b)))
}
