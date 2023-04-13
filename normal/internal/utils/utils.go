package utils

import (
	"chess_normal/internal/entity"
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

func PrintBoard(chess [8][8]bool) {
	alphabet := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for i := 7; i > -1; i-- {
		fmt.Print(i+1, "	")
		for j := 0; j < 8; j++ {
			fmt.Print(chess[i][j], "	")
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Print("	")
	for _, letter := range alphabet {
		fmt.Print("  ", letter, "	")
	}
	fmt.Println()
}

func Convert(position entity.Position) string {
	alphabet := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return alphabet[position.Column] + strconv.Itoa(position.Row+1)
}

func Unique(intSlice []string) []string {
	var list []string
	keys := make(map[string]bool)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func SliceToString(slice []string) string {
	var out string
	for i, str := range slice {
		out += str
		if i < len(slice)-1 {
			out += ","
		}
	}
	return out
}
