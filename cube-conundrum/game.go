package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Draw struct {
	Red   int
	Blue  int
	Green int
}

func (d Draw) String() string {
	return fmt.Sprintf("%d red, %d green, %d blue", d.Red, d.Green, d.Blue)
}

type Game struct {
	Id      int
	Results []Draw
}

func (g Game) String() string {
	var results []string

	for _, res := range results {
		results = append(results, string(res))
	}

	res := strings.Join(results, ";")

	return fmt.Sprintf("Game %d: %s", g.Id, res)
}

func NewGame(line string) (*Game, error) {
	parts := strings.Split(line, ": ")

	// Parse the Game ID
	idstr, have := strings.CutPrefix(parts[0], "Game ")
	if !have {
		return nil, errors.New("did not contain Game prefix")
	}
	gameID, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, err
	}

	// Parse each Draw
	var parsedDraws []Draw
	draws := strings.Split(parts[1], "; ")
	for _, d := range draws {
		var currentDraw Draw

		colorCounts := strings.Split(d, ", ")

		for _, c := range colorCounts {
			var count int
			var color string
			num, err := fmt.Sscanf(c, "%d %s", &count, &color)
			if num != 2 || err != nil {
				return nil, err
			}

			switch color {
			case "red":
				currentDraw.Red = count
			case "green":
				currentDraw.Green = count
			case "blue":
				currentDraw.Blue = count
			}
		}

		parsedDraws = append(parsedDraws, currentDraw)
	}

	return &Game{
		Id:      gameID,
		Results: parsedDraws,
	}, nil
}

func main() {

	result := "dgfa"

	fmt.Printf("Result:\t%v\n", result)
}
