package main

import (
	"aoc23/utils"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Could not read file!")
		return
	}
	defer file.Close()

	partOne(file)

	utils.SeekToFileStart(file)

	partTwo(file)
}

type Coordinate struct {
	X int
	Y int
}

type Dimension struct {
	X int
	Y int
}

const (
	VPipe         byte = 124
	HPipe         byte = 45
	TopToRight    byte = 76
	TopToLeft     byte = 74
	BottomToLeft  byte = 55
	BottomToRight byte = 70
	Ground        byte = 76
)

func partOne(file *os.File) {
	pipeMap := parseMap(file)
	startCoordinate, _ := findStartCoordinates(pipeMap)
	startingDirs := findDirsFromStartingCoordinate(pipeMap, startCoordinate)

	//Replacing S with the inferred point
	pipeMap[startCoordinate.X][startCoordinate.Y] = resolveStartingPointPipeType(startCoordinate, startingDirs)
	var visitedNodes = make(map[Coordinate]int)

	//Walking in one direction away from the start
	walkMap(pipeMap, startCoordinate, startingDirs[0], visitedNodes)

	//Walking in second direction away from the start
	walkMap(pipeMap, startCoordinate, startingDirs[1], visitedNodes)

	var maxSteps = math.MinInt
	for _, steps := range visitedNodes {
		if steps > maxSteps {
			maxSteps = steps
		}
	}

	fmt.Println("Max Part One:", maxSteps)
}

func walkMap(pipeMap [][]byte, start Coordinate, dir Coordinate, nodes map[Coordinate]int) {
	curr := dir
	prev := start
	next := start
	steps := 1
	for curr.X != start.X || curr.Y != start.Y {
		nodeVal, visited := nodes[curr]
		if visited {
			if steps < nodeVal {
				nodes[curr] = steps
			}
		} else {
			nodes[curr] = steps
		}
		steps++

		next, _ = findNextCell(pipeMap, curr, prev)
		prev = curr
		curr = next
	}
}

func findDirsFromStartingCoordinate(pipeMap [][]byte, startingCoordinate Coordinate) []Coordinate {
	var toReturn []Coordinate
	neighbours := findNeighbours(startingCoordinate, Dimension{X: len(pipeMap), Y: len(pipeMap[0])})

	for _, neighbour := range neighbours {
		if neighbour.X == startingCoordinate.X+1 {
			// Below
			if pipeMap[neighbour.X][neighbour.Y] == VPipe || pipeMap[neighbour.X][neighbour.Y] == TopToRight || pipeMap[neighbour.X][neighbour.Y] == TopToLeft {
				toReturn = append(toReturn, neighbour)

			}

		}
		if neighbour.X == startingCoordinate.X-1 {
			// Above
			if pipeMap[neighbour.X][neighbour.Y] == VPipe || pipeMap[neighbour.X][neighbour.Y] == BottomToRight || pipeMap[neighbour.X][neighbour.Y] == BottomToLeft {
				toReturn = append(toReturn, neighbour)
			}
		}
		if neighbour.Y == startingCoordinate.Y+1 {
			// Right
			if pipeMap[neighbour.X][neighbour.Y] == HPipe || pipeMap[neighbour.X][neighbour.Y] == TopToLeft || pipeMap[neighbour.X][neighbour.Y] == BottomToLeft {
				toReturn = append(toReturn, neighbour)

			}

		}
		if neighbour.Y == startingCoordinate.Y-1 {
			// Left
			if pipeMap[neighbour.X][neighbour.Y] == HPipe || pipeMap[neighbour.X][neighbour.Y] == BottomToRight || pipeMap[neighbour.X][neighbour.Y] == TopToRight {
				toReturn = append(toReturn, neighbour)

			}
		}
	}

	return toReturn
}

func findNextCell(pipeMap [][]byte, curr Coordinate, prev Coordinate) (Coordinate, error) {
	switch pipeMap[curr.X][curr.Y] {
	case VPipe:
		return Coordinate{X: 2*curr.X - (prev.X), Y: curr.Y}, nil
	case HPipe:
		return Coordinate{X: curr.X, Y: 2*curr.Y - (prev.Y)}, nil
	case TopToRight:
		switch curr.X - prev.X {
		case 0:
			return Coordinate{X: curr.X - 1, Y: curr.Y}, nil
		case 1:
			return Coordinate{X: curr.X, Y: curr.Y + 1}, nil
		}
	case TopToLeft:
		switch curr.X - prev.X {
		case 0:
			return Coordinate{X: curr.X - 1, Y: curr.Y}, nil
		case 1:
			return Coordinate{X: curr.X, Y: curr.Y - 1}, nil
		}
	case BottomToRight:
		switch curr.X - prev.X {
		case 0:
			return Coordinate{X: curr.X + 1, Y: curr.Y}, nil
		case -1:
			return Coordinate{X: curr.X, Y: curr.Y + 1}, nil
		}
	case BottomToLeft:
		switch curr.X - prev.X {
		case 0:
			return Coordinate{X: curr.X + 1, Y: curr.Y}, nil
		case -1:
			return Coordinate{X: curr.X, Y: curr.Y - 1}, nil
		}
	default:
		return Coordinate{}, errors.New("could not resolvenext cell as no case matched")
	}

	return Coordinate{}, errors.New("could not resolvenext cell as no case matched")
}

func parseMap(file *os.File) [][]byte {
	var pipeMap [][]byte
	scanner := utils.SetupScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		pipeMap = append(pipeMap, line)
	}

	return pipeMap
}

func findStartCoordinates(pipeMap [][]byte) (Coordinate, error) {
	for i := 0; i < len(pipeMap); i++ {
		for j := 0; j < len(pipeMap[0]); j++ {
			if pipeMap[i][j] == 83 {
				return Coordinate{X: i, Y: j}, nil
			}
		}
	}

	return Coordinate{}, errors.New("could not find S in map")
}

func resolveStartingPointPipeType(start Coordinate, dirs []Coordinate) byte {
	switch {
	case isVPipe(start, dirs):
		return VPipe
	case isHPipe(start, dirs):
		return HPipe
	case isTopToRight(start, dirs):
		return TopToRight
	case isTopToLeft(start, dirs):
		return TopToLeft
	case isBottomToRight(start, dirs):
		return BottomToRight
	case isBottomToLeft(start, dirs):
		return BottomToLeft
	default:
		return Ground
	}
}

func isVPipe(start Coordinate, dirs []Coordinate) bool {
	return slices.Contains(dirs, Coordinate{X: start.X + 1, Y: start.Y}) && slices.Contains(dirs, Coordinate{X: start.X - 1, Y: start.Y})
}
func isHPipe(start Coordinate, dirs []Coordinate) bool {
	return slices.Contains(dirs, Coordinate{X: start.X, Y: start.Y + 1}) && slices.Contains(dirs, Coordinate{X: start.X, Y: start.Y - 1})
}
func isTopToRight(start Coordinate, dirs []Coordinate) bool {
	return slices.Contains(dirs, Coordinate{X: start.X - 1, Y: start.Y}) && slices.Contains(dirs, Coordinate{X: start.X, Y: start.Y + 1})
}
func isTopToLeft(start Coordinate, dirs []Coordinate) bool {
	return slices.Contains(dirs, Coordinate{X: start.X - 1, Y: start.Y}) && slices.Contains(dirs, Coordinate{X: start.X, Y: start.Y - 1})
}
func isBottomToRight(start Coordinate, dirs []Coordinate) bool {
	return slices.Contains(dirs, Coordinate{X: start.X + 1, Y: start.Y}) && slices.Contains(dirs, Coordinate{X: start.X, Y: start.Y + 1})
}
func isBottomToLeft(start Coordinate, dirs []Coordinate) bool {
	return slices.Contains(dirs, Coordinate{X: start.X + 1, Y: start.Y}) && slices.Contains(dirs, Coordinate{X: start.X, Y: start.Y - 1})
}

func findNeighbours(cord Coordinate, dim Dimension) []Coordinate {
	var neighbours []Coordinate

	addends := [2]int{1, -1}

	for _, addend := range addends {
		possibleNeighbour := Coordinate{cord.X, cord.Y + addend}
		if isOutOfBound(possibleNeighbour, dim) {
			continue
		}

		neighbours = append(neighbours, Coordinate{cord.X, cord.Y + addend})
	}

	for _, addend := range addends {
		possibleNeighbour := Coordinate{cord.X + addend, cord.Y}
		if isOutOfBound(possibleNeighbour, dim) {
			continue
		}

		neighbours = append(neighbours, Coordinate{cord.X + addend, cord.Y})
	}

	return neighbours
}

func isOutOfBound(coord Coordinate, dim Dimension) bool {
	return coord.X < 0 || coord.X >= dim.X || coord.Y < 0 || coord.Y >= dim.Y
}
