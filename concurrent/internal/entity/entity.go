package entity

var Data []string

var SecondQueen = make(chan Condition, 1)
var ThirdQueen = make(chan Condition, 1)
var FourthQueen = make(chan Condition, 1)
var FifthQueen = make(chan Condition, 1)
var SixthQueen = make(chan Condition, 1)
var SeventhQueen = make(chan Condition, 1)
var EighthQueen = make(chan Condition, 1)

type Position struct {
	Row    int
	Column int
}

type Condition struct {
	Board       [8][8]bool
	Pos         Position
	Coordinates []Position
}
