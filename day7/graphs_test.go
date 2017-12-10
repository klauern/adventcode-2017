package day7

import (
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want program
	}{
		{
			"fwft (72) -> ktlj, cntj, xhth",
			args{"fwft (72) -> ktlj, cntj, xhth"},
			program{"fwft", 72, []string{"ktlj", "cntj", "xhth"}},
		},
		{
			"havc (66)",
			args{"havc (66)"},
			program{"havc", 66, []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []program
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeTree(t *testing.T) {
	type args struct {
		progs []program
	}
	tests := []struct {
		name string
		args args
		want tree
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeTree(tt.args.progs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_hasParents(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name string
		t    tree
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.hasParents(tt.args.program); got != tt.want {
				t.Errorf("tree.hasParents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_getParents(t *testing.T) {
	tests := []struct {
		name string
		t    tree
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.getParents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tree.getParents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_findRoot(t *testing.T) {
	tests := []struct {
		name string
		t    tree
		want program
	}{
		{
			"testAnswer",
			makeTree(parseInput(testInput)),
			makeTree(parseInput(testInput))["tknk"],
		},
		{
			"part 1",
			makeTree(parseInput(getFileString("input.txt"))),
			program{"fbgguv", 57, []string{"bkomu", "tynxlau", "sfruur", "zxvbb", "khupkt", "xkntkvz"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.findRoot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tree.findRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
