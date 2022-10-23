package game

import (
	"errors"
	"fmt"
	"strings"
)

type Room [][]uint8

type Game struct {
	Room        Room
	PositionX   int    `json:"positionX"`
	PositionY   int    `json:"positionY"`
	Orientation string `json:"orientation,omitempty"`
	Command     string `json:"command,omitempty"`
}

func NewGame(x, y, px, py int) *Game {
	var newGame Game
	a := make([][]uint8, x)
	for i := range a {
		a[i] = make([]uint8, y)
	}
	newGame.Room = a
	newGame.Room[px][py] = 1
	newGame.PositionX = px
	newGame.PositionY = py
	return &newGame
}

func UpdateOrientation(game *Game, turnLR string) {
	var newOr string
	if turnLR == "R" {
		switch game.Orientation {
		case "N":
			newOr = "E"
		case "S":
			newOr = "W"
		case "E":
			newOr = "S"
		case "W":
			newOr = "N"
		}
	} else if turnLR == "L" {
		switch game.Orientation {
		case "N":
			newOr = "W"
		case "S":
			newOr = "E"
		case "E":
			newOr = "N"
		case "W":
			newOr = "S"
		}
	}
	game.Orientation = newOr
}

func MoveForward(game *Game) error {

	moveError := errors.New("invalid move")

	switch game.Orientation {
	case "N":
		if game.PositionY == 0 {
			return moveError
		}
		game.Room[game.PositionX][game.PositionY] = 0
		game.Room[game.PositionX][game.PositionY-1] = 1
		game.PositionY--
	case "S":
		if game.PositionY == len(game.Room[0])-1 {
			return moveError
		}

		game.Room[game.PositionX][game.PositionY] = 0
		game.Room[game.PositionX][game.PositionY+1] = 1
		game.PositionY++
	case "E":
		if game.PositionX == len(game.Room)-1 {
			return moveError
		}
		game.Room[game.PositionX][game.PositionY] = 0
		game.Room[game.PositionX+1][game.PositionY] = 1
		game.PositionX++
	case "W":
		if game.PositionX == 0 {
			return moveError
		}
		game.Room[game.PositionX][game.PositionY] = 0
		game.Room[game.PositionX-1][game.PositionY] = 1
		game.PositionX--
	}
	return nil
}

func RunCommand(game *Game) {
	command := game.Command
	cmds := strings.Split(command, "")
	for _, c := range cmds {
		if c == "F" {
			err := MoveForward(game)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			UpdateOrientation(game, c)
		}
	}

}
