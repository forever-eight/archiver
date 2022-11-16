package vlc

import "strings"

type Node struct {
	Value string
	Zero  *Node
	One   *Node
}

func (et encodingTable) DecodingTree() *Node {
	res := &Node{}

	for i, code := range et {
		res.Add(code, i)
	}

	return res
}

func (n *Node) Decode(str string) string {
	var buf strings.Builder

	currentNode := n

	for _, ch := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = n
		}
		switch ch {
		case '0':
			currentNode = currentNode.Zero
		case '1':
			currentNode = currentNode.One
		}
	}

	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
		currentNode = n
	}

	return buf.String()
}

func (n *Node) Add(code string, value rune) {
	// 000101(0) <- value
	currentNode := n
	for _, ch := range code {
		switch ch {
		case '0':
			if currentNode.Zero == nil {
				currentNode.Zero = &Node{}
			}
			currentNode = currentNode.Zero
		case '1':
			if currentNode.One == nil {
				currentNode.One = &Node{}
			}
			currentNode = currentNode.One
		}
	}
	currentNode.Value = string(value)
}
