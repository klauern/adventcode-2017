package day8

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Operation int
type Comparator int

const (
	INC Operation = iota
	DEC
	NONE
	LT Comparator = iota
	LTE
	GT
	GTE
	EQU
	NOT
	NULL
)

type comparison struct {
	register string
	comp     Comparator
	value    int
}

type instruction struct {
	register   string
	operation  Operation
	amount     int
	comparison comparison
}

type registers map[string]int

func newComparator(comp string) Comparator {
	switch comp {
	case "<=":
		return LTE
	case "<":
		return LT
	case ">":
		return GT
	case ">=":
		return GTE
	case "==":
		return EQU
	case "!=":
		return NOT
	default:
		return NULL
	}
}

func newComparison(fields []string) comparison {
	v, err := strconv.Atoi(fields[2])
	if err != nil {
		v = 0
	}
	return comparison{
		register: fields[0],
		comp:     newComparator(fields[1]),
		value:    v,
	}
}

func newOperator(op string) Operation {
	switch op {
	case "dec":
		return DEC
	case "inc":
		return INC
	default:
		return NONE
	}
}

func newInstruction(line string) instruction {
	fields := strings.Fields(line)
	amtBy, err := strconv.Atoi(fields[2])
	if err != nil {
		amtBy = 0
	}
	return instruction{
		register:   fields[0],
		operation:  newOperator(fields[1]),
		amount:     amtBy,
		comparison: newComparison(fields[4:7]),
	}
}

func (r registers) execute(inst instruction) {
	if r.evaluate(inst.comparison) {
		switch inst.operation {
		case DEC:
			r[inst.register] -= inst.amount
		case INC:
			r[inst.register] += inst.amount
		default:
			break
		}
	}
}

func (r registers) evaluate(comp comparison) bool {
	reg := r[comp.register]
	switch comp.comp {
	case LTE:
		return reg <= comp.value
	case LT:
		return reg < comp.value
	case GT:
		return reg > comp.value
	case GTE:
		return reg >= comp.value
	case EQU:
		return reg == comp.value
	case NOT:
		return reg != comp.value
	default:
		return false
	}
}

func createRegister(input string) registers {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instructions[i] = newInstruction(line)
	}

	reg := make(registers)
	for _, instruction := range instructions {
		reg.execute(instruction)
	}

	return reg
}

func createRegisterWithFile() registers {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return registers{}
	}
	return createRegister(string(file))
}
