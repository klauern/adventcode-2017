package day9

type Garbage int
type Group int

const (
	LBRACE Garbage = iota
	RBRACE
	LT Group = iota
	GT
)
