package game

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	fmt.Println("test if a new game is corectly generated")
	game := NewGame(2, 3, 1, 2)
	if len(game.Room) != 2 {
		t.Errorf("UpdateOrientation failed. Expected: %d, Got: %d", 2, len(game.Room))
	}
	if len(game.Room[0]) != 3 {
		t.Errorf("UpdateOrientation failed. Expected: %d, Got: %d", 3, len(game.Room[0]))
	}
}

func TestUpdateOrientation(t *testing.T) {
	fmt.Println("test if the orientation is correctly updated")
	game := NewGame(5, 5, 1, 2)
	game.Orientation = "W"
	game.Command = ""
	fmt.Println("testing")

	UpdateOrientation(game, "L")
	if game.Orientation != "S" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "S", game.Orientation)
	}

	UpdateOrientation(game, "L")
	if game.Orientation != "E" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "E", game.Orientation)
	}

	UpdateOrientation(game, "L")
	if game.Orientation != "N" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "N", game.Orientation)
	}

	UpdateOrientation(game, "L")
	if game.Orientation != "W" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "W", game.Orientation)
	}

	UpdateOrientation(game, "R")
	if game.Orientation != "N" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "N", game.Orientation)
	}

	UpdateOrientation(game, "R")
	if game.Orientation != "E" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "E", game.Orientation)
	}

	UpdateOrientation(game, "R")
	if game.Orientation != "S" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "S", game.Orientation)
	}

	UpdateOrientation(game, "R")
	if game.Orientation != "W" {
		t.Errorf("UpdateOrientation failed. Expected: %s, Got: %s", "W", game.Orientation)
	}
}

func TestMoveForward(t *testing.T) {
	fmt.Println("respond with 'invalid move' if the robot try to walk over the room border")
	game := NewGame(7, 3, 6, 2)
	game.Command = "F"

	game.Orientation = "S"
	err := MoveForward(game)
	if err == nil {
		t.Errorf("TestMoveForward failed. Expected: %v, Got: %v", errors.New("invalid move"), err)
	}

	game.Orientation = "N"
	err = MoveForward(game)
	if err != nil {
		t.Errorf("TestMoveForward failed")
	}

	game.Orientation = "E"
	err = MoveForward(game)
	if err == nil {
		t.Errorf("TestMoveForward failed")
	}

	game.Orientation = "W"
	err = MoveForward(game)
	if err != nil {
		t.Errorf("TestMoveForward failed")
	}
}

func TestRunCommand(t *testing.T) {
	type result struct {
		x int
		y int
		o string
	}
	fmt.Println("test the robot normal walk")
	var commands = []string{"FFFR", "FRFFFLFL", "FRF", "RFFLFRFR"}
	var results = []result{{3, 0, "S"}, {2, 3, "N"}, {1, 1, "S"}, {1, 3, "W"}}
	for i, cmd := range commands {
		game := NewGame(4, 6, 0, 0)
		game.Orientation = "E"
		game.Command = cmd
		RunCommand(game)

		r := result{game.PositionX, game.PositionY, game.Orientation}
		if !reflect.DeepEqual(r, results[i]) {
			t.Errorf("Incorrect result, expected: %v, get: %v", results[i], r)
		}

	}

}
func TestRunCommand1(t *testing.T) {
	type result struct {
		x int
		y int
		o string
	}

	// test that the robot can walk continuosly
	fmt.Println("test that the robot can walk continuosly")
	game1 := NewGame(4, 6, 0, 0)
	game1.Orientation = "E"
	var commands1 = []string{"FFFR", "FFFFFL", "FRFFFFLFFL", "FLFRFFFRFR"}
	var results1 = []result{{3, 0, "S"}, {3, 5, "W"}, {0, 1, "S"}, {0, 5, "N"}}
	for i, cmd1 := range commands1 {
		game1.Command = cmd1
		RunCommand(game1)
		game1.Orientation = results1[i].o
		game1.PositionX = results1[i].x
		game1.PositionY = results1[i].y

		r1 := result{game1.PositionX, game1.PositionY, game1.Orientation}
		if !reflect.DeepEqual(r1, results1[i]) {
			t.Errorf("Incorrect result, expected: %v, get: %v", results1[i], r1)
		}
	}

}
