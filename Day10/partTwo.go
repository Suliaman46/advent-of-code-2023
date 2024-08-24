package main

import (
	"fmt"
	"os"
)

func partTwo(file *os.File) {
	pipeMap := parseMap(file)
	startCoordinate, _ := findStartCoordinates(pipeMap)
	startingDirs := findDirsFromStartingCoordinate(pipeMap, startCoordinate)

	//Replacing S with the inferred point
	pipeMap[startCoordinate.X][startCoordinate.Y] = resolveStartingPointPipeType(startCoordinate, startingDirs)
	var pathNodes = make(map[Coordinate]int)

	//Walking in one direction away from the start
	walkMap(pipeMap, startCoordinate, startingDirs[0], pathNodes)

	//Add Start to Visited Nodes
	pathNodes[startCoordinate] = 0

	isInside := false
	count := 0
	for i := 0; i < len(pipeMap); i++ {
		for j := 0; j < len(pipeMap[0]); j++ {
			_, onPath := pathNodes[Coordinate{i, j}]
			if onPath {
				pipeCell := pipeMap[i][j]
				switch pipeCell {
				case VPipe:
					fallthrough
				case TopToLeft:
					fallthrough
				case TopToRight:
					isInside = !isInside
				}
			} else {
				if isInside {
					count++

				}
			}

		}
		isInside = false
	}

	fmt.Println("Total inside Part Two: ", count)
}
