package main

import (
	"fmt"
	"strings"
)

type BatleField struct {
	Size int
}

func (p *BatleField) StringRender(x int) {
	p.Size = x
	GameMap := CreateEmptyMap(p.Size)
	for i := 0; i < p.Size+2; i++ {
		out := strings.Join(GameMap[i], " ")
		fmt.Println(out)
	}
}

func main() {
	var field BatleField
	fmt.Scan(&field.Size)
	field.StringRender(field.Size)
}

func CreateEmptyMap(x int) [][]string {
	x += 2 // Это пока для наглядности
	GameMap := make([][]string, x)
	for i := range GameMap {
		GameMap[i] = make([]string, x)
	}
	for i := range GameMap {
		if i == 0 || i == x-1 {
			for j := range GameMap[i] {
				GameMap[i][j] = "0"
			}
		} else {
			GameMap[i][0] = "0"
			GameMap[i][x-1] = "0"
			for j := 1; j < x-1; j++ {
				GameMap[i][j] = "-"
			}
		}
	}
	return GameMap
}
