package main

import (
	"fmt"
	"strings"
)

type BattleField struct {
	Size int
}

func (p *BattleField) StringRender() {
	GameMap := CreateEmptyMap(p.Size)
	for i := 0; i < p.Size+2; i++ {
		out := strings.Join(GameMap[i], " ")
		fmt.Println(out)
	}
}

const (
	border = "0"
	cell   = "-"
)

func main() {
	var field BattleField
	fmt.Scan(&field.Size)
	field.StringRender()
}

func CreateEmptyMap(x int) [][]string {
	GameMap := make([][]string, x+2)
	for i := range GameMap {
		GameMap[i] = make([]string, x+2)
	}
	for i := range GameMap {
		if i == 0 || i == x+1 {
			for j := range GameMap[i] {
				GameMap[i][j] = border
			}
		} else {
			GameMap[i][0] = border
			GameMap[i][x+1] = border
			for j := 1; j < x+1; j++ {
				GameMap[i][j] = cell
			}
		}
	}
	return GameMap
}
