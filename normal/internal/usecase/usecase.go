package usecase

import (
	"chess_normal/internal/entity"
	"chess_normal/internal/utils"
	"sort"
)

func Cal() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			var board [8][8]bool
			p := entity.Position{
				Row:    i,
				Column: j,
			}
			c := entity.Condition{
				Pos:   p,
				Board: board,
			}
			c = utils.Unavailable(c)
			check(c)
		}
	}
}

func check(c entity.Condition) {
	var test bool
	for rKey, row := range c.Board {
		for cKey, cell := range row {
			if !cell {
				test = true
				p := entity.Position{
					Row:    rKey,
					Column: cKey,
				}
				c.Pos = p
				addQueen(c)
				break
			}
		}
	}
	if !test {
		if len(c.Coordinates) > 7 {
			var slice []string
			for _, coord := range c.Coordinates {
				slice = append(slice, utils.Convert(coord))
			}
			sort.Strings(slice)
			entity.Data = append(entity.Data, utils.SliceToString(slice))
		}
	}
}

func addQueen(c entity.Condition) {
	c = utils.Unavailable(c)
	//printBoard(c.board)
	check(c)
}
