package usecase

import (
	"ChessConcurrent/internal/entity"
	"ChessConcurrent/internal/utils"
	"sync"
)

func AddFirst(wg *sync.WaitGroup) {
	defer wg.Done()
	var b [8][8]bool

	c := entity.Condition{
		Board: b,
	}

	for j := 0; j < 8; j++ {
		p := entity.Position{
			Row:    0,
			Column: j,
		}
		c.Pos = p
		entity.SecondQueen <- utils.Unavailable(c)
	}
	close(entity.SecondQueen)
}

func AddSecond(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range entity.SecondQueen {
		for j := 0; j < 8; j++ {
			if !c.Board[1][j] {
				p := entity.Position{
					Row:    1,
					Column: j,
				}
				c.Pos = p
				entity.ThirdQueen <- utils.Unavailable(c)
			}
		}
	}
	close(entity.ThirdQueen)
}

func AddThird(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range entity.ThirdQueen {
		for j := 0; j < 8; j++ {
			if !c.Board[2][j] {
				p := entity.Position{
					Row:    2,
					Column: j,
				}
				c.Pos = p
				entity.FourthQueen <- utils.Unavailable(c)
			}
		}
	}
	close(entity.FourthQueen)
}

func AddFourth(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range entity.FourthQueen {
		for j := 0; j < 8; j++ {
			if !c.Board[3][j] {
				p := entity.Position{
					Row:    3,
					Column: j,
				}
				c.Pos = p
				entity.FifthQueen <- utils.Unavailable(c)
			}
		}
	}
	close(entity.FifthQueen)
}

func AddFifth(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range entity.FifthQueen {
		for j := 0; j < 8; j++ {
			if !c.Board[4][j] {
				p := entity.Position{
					Row:    4,
					Column: j,
				}
				c.Pos = p
				entity.SixthQueen <- utils.Unavailable(c)
			}
		}
	}
	close(entity.SixthQueen)
}

func AddSixth(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range entity.SixthQueen {
		for j := 0; j < 8; j++ {
			if !c.Board[5][j] {
				p := entity.Position{
					Row:    5,
					Column: j,
				}
				c.Pos = p
				entity.SeventhQueen <- utils.Unavailable(c)
			}
		}
	}
	close(entity.SeventhQueen)
}

func AddSeventh(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range entity.SeventhQueen {
		for j := 0; j < 8; j++ {
			if !c.Board[6][j] {
				p := entity.Position{
					Row:    6,
					Column: j,
				}
				c.Pos = p
				entity.EighthQueen <- utils.Unavailable(c)
			}
		}
	}
	close(entity.EighthQueen)
}

func AddEighth(wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range entity.EighthQueen {
		for j := 0; j < 8; j++ {
			if !c.Board[7][j] {
				var str string
				p := entity.Position{
					Row:    7,
					Column: j,
				}
				c.Pos = p
				c = utils.Unavailable(c)
				for index, coord := range c.Coordinates {
					str += utils.Convert(coord)
					if index < 7 {
						str += ","
					}
				}
				entity.Data = append(entity.Data, str)
			}
		}
	}
}
