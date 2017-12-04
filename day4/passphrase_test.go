package day4

import (
	"reflect"
	"testing"
)

func Test_mapPassphrase(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name string
		args args
		want passphrase
	}{
		{
			"aa bb cc dd ee",
			args{"aa bb cc dd ee"},
			passphrase{"aa": 1, "bb": 1, "cc": 1, "dd": 1, "ee": 1},
		},
		{
			"aa bb cc dd aaa",
			args{"aa bb cc dd aaa"},
			passphrase{"aa": 1, "bb": 1, "cc": 1, "dd": 1, "aaa": 1},
		},
		{
			"aa bb cc dd aa",
			args{"aa bb cc dd aa"},
			passphrase{"aa": 2, "bb": 1, "cc": 1, "dd": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapPassphrase(tt.args.phrase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapPassphrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidPassPhrase(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"aa bb cc dd ee",
			args{"aa bb cc dd ee"},
			true,
		},
		{
			"aa bb cc dd aa",
			args{"aa bb cc dd aa"},
			false,
		},
		{
			"aa bb cc dd aaa",
			args{"aa bb cc dd aaa"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPassPhrase(tt.args.phrase); got != tt.want {
				t.Errorf("ValidPassPhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart1Test(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"input file",
			477,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1Test(); got != tt.want {
				t.Errorf("Part1Test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidPassPhraseNoAnagrams(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcde fghij",
			args{"abcde fghij"},
			true,
		},
		{
			"abcde xyz ecdab",
			args{"abcde xyz ecdab"},
			false,
		},
		{
			"a ab abc abd abf abj",
			args{"a ab abc abd abf abj"},
			true,
		},
		{
			"iiii oiii ooii oooi oooo",
			args{"iiii oiii ooii oooi oooo"},
			true,
		},
		{
			"oiii ioii iioi iiio",
			args{"oiii ioii iioi iiio"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPassPhraseNoAnagrams(tt.args.str); got != tt.want {
				t.Errorf("ValidPassPhraseNoAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2Test(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"input.txt",
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2Test(); got != tt.want {
				t.Errorf("Part2Test() = %v, want %v", got, tt.want)
			}
		})
	}
}
