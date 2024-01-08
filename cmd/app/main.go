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
	Player Player
}

type Player struct {
	cordX int
	cordY int
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
	i, j := p.Player.cordY, p.Player.cordX
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
			return
		}
		switch string(char) {
		case "w":
			if p.Player.cordY != 0 {
				p.Player.cordY--
			}
		case "a":
			if p.Player.cordX != 0 {
				p.Player.cordX--
			}
		case "s":
			if p.Player.cordY != p.Size-1 {
				p.Player.cordY++
			}
		case "d":
			if p.Player.cordX != p.Size-1 {
				p.Player.cordX++
			}

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
	fmt.Scan(&field.Player.cordX, &field.Player.cordY)
	fmt.Println("Press ESC to quit")
	fmt.Println(field)
	field.Moving()
}
