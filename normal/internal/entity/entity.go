package entity

var Data []string

type Position struct {
	Row    int
	Column int
}

type Condition struct {
	Board       [8][8]bool
	Pos         Position
	Coordinates []Position
}
