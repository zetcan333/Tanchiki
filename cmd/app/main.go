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
	out = out[:len(out)-1] + "\n"
	for i := 0; i < p.Size; i++ {
		out += "s "
		out += strings.Join(GameMap[i], " ")
		out += " s\n"
	}
	out += strings.Repeat("s ", p.Size+2)
	out = out[:len(out)-1] + "\n"
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
	fmt.Println(field.String())
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
