package day8

import (
	"reflect"
	"testing"
)

const testInput = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func Test_newComparator(t *testing.T) {
	type args struct {
		comp string
	}
	tests := []struct {
		name string
		args args
		want Comparator
	}{
		{
			"less than",
			args{"<"},
			LT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newComparator(tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newComparison(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name string
		args args
		want comparison
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newComparison(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newComparison() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newOperator(t *testing.T) {
	type args struct {
		op string
	}
	tests := []struct {
		name string
		args args
		want Operation
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newOperator(tt.args.op); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newInstruction(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want instruction
	}{
		{
			"b inc 5 if a > 1",
			args{"b inc 5 if a > 1"},
			instruction{
				amount: 5,
				comparison: comparison{
					comp:     GT,
					register: "a",
					value:    1,
				},
				operation: INC,
				register:  "b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInstruction(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registers_execute(t *testing.T) {
	type args struct {
		inst instruction
	}
	tests := []struct {
		name string
		r    registers
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.execute(tt.args.inst)
		})
	}
}

func Test_registers_evaluate(t *testing.T) {
	type args struct {
		comp comparison
	}
	tests := []struct {
		name string
		r    registers
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.evaluate(tt.args.comp); got != tt.want {
				t.Errorf("registers.evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createRegister(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want registers
	}{
		{
			"input",
			args{testInput},
			registers{
				"a": 1,
				"c": -10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createRegister(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createRegisterWithFile(t *testing.T) {
	tests := []struct {
		name string
		want registers
	}{
		{
			"test input",
			registers{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createRegisterWithFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createRegisterWithFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
