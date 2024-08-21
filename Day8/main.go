package main

import (
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

const StartPosition = "AAA"
const EndPosition = "ZZZ"
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
}

func partOne(file *os.File) {
	desertMap := parseTree(file)
	steps := walkMap(desertMap)
	fmt.Println("Total Steps Part One: ", steps)
}

func parseTree(file *os.File) DesertMap {
	var head *Node

	var seenNodes = make(map[string]*Node)

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

			if pos == StartPosition {
				head = &curr
			}
		}
	}

	return DesertMap{instructions, head}
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

func walkMap(desertMap DesertMap) int {
	var curr *Node
	curr = desertMap.startingNode
	steps := 0
	for i := 0; i < len(desertMap.instructions); i++ {
		instruction := desertMap.instructions[i]

		if curr.pos == EndPosition {
			return steps
		}

		// Reached end of instructions but not ZZZ so restart
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
