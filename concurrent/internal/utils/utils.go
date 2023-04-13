package utils

import (
	"ChessConcurrent/internal/entity"
	"fmt"
	"strconv"
)

func Unavailable(con entity.Condition) entity.Condition {
	var lOrigin, rOrigin [2]int

	//find left origin
	if con.Pos.Row > con.Pos.Column {
		lOrigin = [2]int{con.Pos.Row - con.Pos.Column, 0}
	} else if con.Pos.Column > con.Pos.Row {
		lOrigin = [2]int{0, con.Pos.Column - con.Pos.Row}
	}

	//find right origin
	if con.Pos.Row+con.Pos.Column > 7 {
		rOrigin = [2]int{con.Pos.Row + con.Pos.Column - 7, 7}
	} else {
		rOrigin = [2]int{0, con.Pos.Row + con.Pos.Column}
	}

	for i := 0; i < 8; i++ {
		con.Board[con.Pos.Row][i] = true
		con.Board[i][con.Pos.Column] = true
	}
	con.Board = diagonal(con.Board, lOrigin, rOrigin)

	out := entity.Condition{
		Pos:         con.Pos,
		Board:       con.Board,
		Coordinates: append(con.Coordinates, con.Pos),
	}

	return out
}

func diagonal(chess [8][8]bool, lOrigin, rOrigin [2]int) [8][8]bool {
	for i := 0; i < 8; i++ {
		if lOrigin[0]+i < 8 && lOrigin[1]+i < 8 {
			chess[lOrigin[0]+i][lOrigin[1]+i] = true
		}
		if rOrigin[1]-i > -1 && rOrigin[0]+i < 8 {
			chess[rOrigin[0]+i][rOrigin[1]-i] = true
		}
	}
	return chess
}

func Convert(position entity.Position) string {
	alphabet := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return alphabet[position.Column] + strconv.Itoa(position.Row+1)
}

func PrintConditions() {
	for _, d := range entity.Data {
		fmt.Println(d)
	}
	fmt.Printf("there's %v conditions. ", len(entity.Data))
}
