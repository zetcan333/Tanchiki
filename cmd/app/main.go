package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/eiannone/keyboard"
)

type BattleField struct {
	Size   int
	Player [2]int
}

func (p BattleField) EmptyMap() [][]string {
	Map := make([][]string, p.Size)
	for i := range Map {
		Map[i] = make([]string, p.Size)
		for j := range Map[i] {
			Map[i][j] = cell
		}
	}
	return Map
}

func (p BattleField) AddBorders(Map *[][]string) {
	for i := range *Map {
		(*Map)[i] = append([]string{border}, (*Map)[i]...)
		(*Map)[i] = append((*Map)[i], border)
	}

	borderline := make([][]string, 1)
	borderline[0] = make([]string, p.Size+2)
	for i := range borderline[0] {
		borderline[0][i] = border
	}

	*Map = append([][]string{borderline[0]}, *Map...)
	*Map = append(*Map, borderline[0])
}

func (p BattleField) AddPlayer(GameMap *[][]string) {
	i, j := p.Player[0], p.Player[1]
	(*GameMap)[i][j] = player
}

func (p BattleField) GenerateMap() [][]string {
	Map := p.EmptyMap()
	p.AddPlayer(&Map)
	p.AddBorders(&Map)
	return Map
}

func (p BattleField) String() string {
	Map := p.GenerateMap()
	out := ""

	for i := range Map {
		out += strings.Join(Map[i], " ")
		out += "\n"
	}

	return fmt.Sprint(out)
}

func (p *BattleField) Moving() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		switch string(char) {
		case "w":
			if p.Player[0] == 0 {
				break
			}
			p.Player[0]--
		case "a":
			if p.Player[1] == 0 {
				break
			}
			p.Player[1]--
		case "s":
			if p.Player[0] == p.Size-1 {
				break
			}
			p.Player[0]++
		case "d":
			if p.Player[1] == p.Size-1 {
				break
			}
			p.Player[1]++

		default:
			continue
		}
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Press ESC to quit")
		fmt.Println(p)
	}
}

const (
	border = "0"
	cell   = "-"
	player = "@"
)

func main() {
	var field BattleField
	fmt.Print("Размер поля: ")
	fmt.Scanln(&field.Size)
	fmt.Print("Начальные координаты игрока (x, y): ")
	fmt.Scan(&field.Player[1], &field.Player[0])
	fmt.Println("Press ESC to quit")
	fmt.Println(field)
	field.Moving()
}
