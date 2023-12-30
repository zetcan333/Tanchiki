package main

import (
	"fmt"
	"strings"
)

type BattleField struct {
	Size int
}

func (p *BattleField) String() string {
	GameMap := CreateEmptyMap(p.Size)
	out := strings.Repeat("s ", p.Size+2)
	out, _ = strings.CutSuffix(out, " ")
	for i := 0; i < p.Size; i++ {
		out += "s "
		out += strings.Join(GameMap[i], " ")
		out += " s"
	}
	out += strings.Repeat("s ", p.Size+2)
	out, _ = strings.CutSuffix(out, " ")
	out = strings.Replace(out, "s", border, -1)
	return out
}

const (
	border = "0"
	cell   = "-"
)

func main() {
	var field BattleField
	fmt.Scan(&field.Size)
	render := field.String()
	line := (field.Size+2)*2 - 1
	for i := 0; i < (line * (field.Size + 2)); i += line {
		fmt.Println(render[i : i+line])
	}
}

func CreateEmptyMap(x int) [][]string {
	GameMap := make([][]string, x)
	for i := range GameMap {
		GameMap[i] = make([]string, x)
		for j := range GameMap[i] {
			GameMap[i][j] = cell
		}
	}
	return GameMap
}
