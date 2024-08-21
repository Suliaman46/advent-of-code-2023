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
	parseTree(file)
}

func parseTree(file *os.File) *Node {
	var head *Node

	var seenNodes = make(map[string]*Node)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

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

			if pos == "AAA" {
				head = &curr
			}
		}
	}

	fmt.Println("Head", head)
	for pos, node := range seenNodes {
		if node.right == nil || node.right == nil {
			fmt.Println("Node with nil children at: ", pos)
		}
	}
	return head
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
