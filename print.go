package main

import (
	"fmt"
)

func (bst *Bst[T]) Print() {
	lines := bst.nodeToString(bst.root)
	for _, line := range lines {
		fmt.Println(line)
	}
}

func (bst *Bst[T]) nodeToString(node *Node[T]) []string {
	if node == nil {
		return []string{}
	}
	nodeStr := fmt.Sprintf("%v", node.value)
	if node.left == nil && node.right == nil {
		return []string{nodeStr}
	}
	leftLines := bst.nodeToString(node.left)
	rightLines := bst.nodeToString(node.right)
	leftWidth := widthOfLines(leftLines)
	rightWidth := widthOfLines(rightLines)
	rootWidth := len(nodeStr)
	var lines []string
	if node.left != nil && node.right != nil {
		rootPos := leftWidth
		rootLine := spaces(rootPos) + nodeStr + spaces(rightWidth)
		lines = append(lines, rootLine)
		branchLine := spaces(leftWidth) + "/" + spaces(rootWidth-1) + "\\"
		lines = append(lines, branchLine)
		merged := mergeLines(leftLines, rightLines)
		lines = append(lines, merged...)
		return lines
	} else if node.left != nil {
		rootPos := leftWidth
		rootLine := spaces(rootPos) + nodeStr
		lines = append(lines, rootLine)
		branchLine := spaces(rootPos) + "|"
		lines = append(lines, branchLine)
		lines = append(lines, leftLines...)
		return lines
	} else {
		rootLine := nodeStr
		lines = append(lines, rootLine)
		branchLine := spaces(len(nodeStr)) + "\\"
		lines = append(lines, branchLine)
		indentedRight := indentLines(rightLines, len(nodeStr)+1)
		lines = append(lines, indentedRight...)
		return lines
	}
}

func widthOfLines(lines []string) int {
	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}

func spaces(n int) string {
	if n < 0 {
		return ""
	}
	return fmt.Sprintf("%*s", n, "")
}

func mergeLines(left, right []string) []string {
	lw := widthOfLines(left)
	height := len(left)
	if len(right) > height {
		height = len(right)
	}
	var merged []string
	for i := 0; i < height; i++ {
		var leftLine, rightLine string
		if i < len(left) {
			leftLine = left[i]
		}
		if i < len(right) {
			rightLine = right[i]
		}
		leftLine = leftLine + spaces(lw-len(leftLine))
		line := leftLine + " " + rightLine
		merged = append(merged, line)
	}
	return merged
}

func indentLines(lines []string, indent int) []string {
	var result []string
	sp := spaces(indent)
	for _, l := range lines {
		result = append(result, sp+l)
	}
	return result
}
