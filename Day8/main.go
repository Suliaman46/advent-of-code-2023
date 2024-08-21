package main

import (
	"aoc23/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	pos   string
	left  *Node
	right *Node
}

type DesertMap struct {
	instructions []string
	startingNode *Node
}

type DesertMapPartTwo struct {
	instructions  []string
	startingNodes []*Node
}

const StartPosition = "AAA"
const StartSuffix = "A"
const EndPosition = "ZZZ"
const EndSuffix = "Z"
const LeftStep = "L"
const RightStep = "R"

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

func partOne(file *os.File) {
	desertMap := parseTree(file, func(pos string) bool {
		return pos == StartPosition
	})
	steps := resolveStepsToEnd(desertMap, func(pos string) bool {
		return pos == EndPosition
	})
	fmt.Println("Total Steps Part One: ", steps[0])
}

func partTwo(file *os.File) {
	desertMap := parseTree(file, func(pos string) bool {
		return strings.HasSuffix(pos, StartSuffix)
	})
	steps := resolveStepsToEnd(desertMap, func(pos string) bool {
		return strings.HasSuffix(pos, EndSuffix)
	})
	results := resolveCommonStep(steps)
	fmt.Println("Total Steps Part Two:", results)
}

func parseTree(file *os.File, startIdentifier func(pos string) bool) DesertMapPartTwo {
	var seenNodes = make(map[string]*Node)
	var startingPoints []*Node
	scanner := bufio.NewScanner(file)
	idx := 0
	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()

		if idx == 0 {
			for _, char := range line {
				instructions = append(instructions, string(char))
			}
			idx++
		}

		if !strings.Contains(line, "=") {
			continue
		}

		var pos = line[0:3]
		var posLeft = line[7:10]
		var posRight = line[12:15]

		node, exists := seenNodes[pos]

		if exists {
			handleSeenNode(node, posLeft, posRight, &seenNodes)
		} else {

			curr := Node{pos, nil, nil}

			handleLeftChild(&curr, posLeft, &seenNodes)
			handleRightChild(&curr, posRight, &seenNodes)

			seenNodes[pos] = &curr

			if startIdentifier(pos) {
				startingPoints = append(startingPoints, &curr)
			}
		}
	}

	return DesertMapPartTwo{instructions, startingPoints}
}

func handleSeenNode(node *Node, posLeft string, posRight string, seenNodes *map[string]*Node) {
	handleLeftChildForSeenNode(node, posLeft, seenNodes)
	handleRightChildForSeenNode(node, posRight, seenNodes)
}

func handleLeftChildForSeenNode(node *Node, posLeft string, seenNodes *map[string]*Node) {
	if node.left != nil {
		return
	}

	if node.pos == posLeft {
		node.left = node
		return
	}

	handleLeftChild(node, posLeft, seenNodes)
}

func handleRightChildForSeenNode(node *Node, posRight string, seenNodes *map[string]*Node) {
	if node.right != nil {
		return
	}

	if node.pos == posRight {
		node.right = node
	}

	handleRightChild(node, posRight, seenNodes)
}

func handleLeftChild(node *Node, posLeft string, seenNodes *map[string]*Node) {
	left, leftExists := (*seenNodes)[posLeft]
	if leftExists {

		node.left = left
		return
	}

	newLeft := Node{posLeft, nil, nil}
	node.left = &newLeft
	(*seenNodes)[posLeft] = &newLeft
}

func handleRightChild(node *Node, posRight string, seenNodes *map[string]*Node) {
	right, rightExists := (*seenNodes)[posRight]
	if rightExists {
		node.right = right
		return
	}

	newChild := Node{posRight, nil, nil}
	node.right = &newChild
	(*seenNodes)[posRight] = &newChild
}

func resolveStepsToEnd(desertMap DesertMapPartTwo, endIdentifier func(pos string) bool) []int {
	startingNodes := desertMap.startingNodes
	steps := 0

	stepsTaken := make([]int, len(startingNodes))
	for i := 0; i < len(startingNodes); i++ {
		startingNode := startingNodes[i]
		steps = walkMap(DesertMap{desertMap.instructions, startingNode}, endIdentifier)

		stepsTaken[i] = steps
	}

	return stepsTaken
}

func walkMap(desertMap DesertMap, endIdentifier func(pos string) bool) int {
	var curr *Node
	curr = desertMap.startingNode
	steps := 0
	for i := 0; i < len(desertMap.instructions); i++ {
		instruction := desertMap.instructions[i]

		if endIdentifier(curr.pos) {
			return steps
		}

		// Reached end of instructions but not END node so restart
		if i == len(desertMap.instructions)-1 {
			i = -1
		}

		switch instruction {
		case LeftStep:
			curr = curr.left
		case RightStep:
			curr = curr.right
		}

		steps++
	}

	return steps
}

func resolveCommonStep(steps []int) int {
	result := steps[0]
	for i := 1; i < len(steps); i++ {
		result = LCM(result, steps[i])
	}

	return result
}

func LCM(a, b int) int {
	result := a * b / GCD(a, b)

	return result
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
