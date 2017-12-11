package day7

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type program struct {
	name     string
	weight   int
	children []string
}

type tree map[string]program

const testInput = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func getFileString(input string) string {
	file, err := ioutil.ReadFile(input)
	if err != nil {
		return ""
	}
	return string(file)
}

func parseInput(input string) []program {
	lines := strings.Split(input, "\n")
	programs := make([]program, len(lines))
	for i, line := range lines {
		programs[i] = parseLine(line)
	}
	return programs
}

func parseLine(line string) program {
	pieces := strings.Split(line, " (")
	name := pieces[0]
	weight := -1
	children := make([]string, 0)
	if strings.Contains(pieces[1], "->") {
		weightRest := strings.Split(pieces[1], ") -> ")
		w, err := strconv.Atoi(weightRest[0])
		if err == nil {
			weight = w
		}
		others := strings.Split(weightRest[1], ", ")
		children = others
	} else {
		wString := strings.TrimSuffix(pieces[1], ")")
		w, err := strconv.Atoi(wString)
		if err == nil {
			weight = w
		}
	}
	return program{name, weight, children}
}

func makeTree(progs []program) tree {
	t := make(tree)
	for _, program := range progs {
		t[program.name] = program
	}
	return t
}

func (t tree) hasParents(program string) bool {
	for leaf, prog := range t {
		if leaf != program {
			for _, v := range prog.children {
				if program == v {
					return true
				}
			}
		}
	}
	return false
}

func (t tree) getParents() []string {
	programs := make([]string, 0)
	for _, prog := range t {
		if len(prog.children) > 0 {
			programs = append(programs, prog.name)
		}
	}
	return programs
}

func (t tree) findRoot() program {
	parents := t.getParents()
	for _, prog := range parents {
		if !t.hasParents(prog) {
			return t[prog]
		}
	}
	return program{}
}

func (t tree) calcTotalWeight(prog program) int {
	weight := prog.weight
	for _, v := range prog.children {
		weight += t.calcTotalWeight(t[v])
	}
	return weight
}

func (t tree) isBalanced(prog program) bool {
	weight := t[prog.children[0]].weight
	for i := 1; i < len(prog.children); i++ {
		if weight != t[prog.children[i]].weight {
			return false
		}
	}
	return true
}

func (t tree) getUnbalancedImmediateChild(prog program) program {
	for _, name := range prog.children {

	}
}
