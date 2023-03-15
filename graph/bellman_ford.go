package graph

import (
	"fmt"
)

type Node struct {
	rate float64
	path []int
}

func BellmanFordWithoutNegativeCycle(edge [][]float64, from int, to int) (rate float64, path []int, err error) {
	nodeNum := len(edge)
	node := make([]Node, nodeNum)
	rate = 0
	err = nil
	path = make([]int, 0)

	for i := 0; i < nodeNum; i++ {
		node[i].rate = -1
	}

	node[from].rate = 1
	node[from].path = []int{from}

	for try := 1; try < nodeNum; try++ {
		for i := 0; i < nodeNum; i++ {
			fromNode := i
			for j := 0; j < nodeNum; j++ {
				toNode := j
				if edge[fromNode][toNode] != -1 && node[fromNode].rate != -1 {
					if node[toNode].rate < node[fromNode].rate*edge[fromNode][toNode] && !isVisited(node[fromNode].path, toNode) {
						node[toNode].rate = node[fromNode].rate * edge[fromNode][toNode]
						node[toNode].path = append(node[fromNode].path, toNode)
					}
				}
			}
		}
	}

	if node[to].rate == -1 {
		err = fmt.Errorf("path was not found. from: %d, to: %d", from, to)
		return
	}
	path = node[to].path
	rate = node[to].rate
	return
}

func BellmanFordWithNegativeCycle(edge [][]float64, from int, to int) (rate float64, path []int, err error) {
	nodeNum := len(edge)
	node := make([]Node, nodeNum)
	rate = 0
	err = nil
	path = make([]int, 0)

	for i := 0; i < nodeNum; i++ {
		node[i].rate = -1
	}

	node[from].rate = 1
	node[from].path = []int{from}

	for try := 1; try <= nodeNum; try++ {
		for i := 0; i < nodeNum; i++ {
			fromNode := i
			for j := 0; j < nodeNum; j++ {
				toNode := j
				if edge[fromNode][toNode] != -1 && node[fromNode].rate != -1 {
					if node[toNode].rate < node[fromNode].rate*edge[fromNode][toNode] {
						if try == nodeNum {
							err = fmt.Errorf("negative cycle %d -> %d", fromNode, toNode)
							path = node[to].path
							rate = node[to].rate
							return
						}
						node[toNode].rate = node[fromNode].rate * edge[fromNode][toNode]
						node[toNode].path = append(node[fromNode].path, toNode)
					}
				}
			}
		}
	}

	if node[to].rate == -1 {
		err = fmt.Errorf("path was not found. from: %d, to: %d", from, to)
		return
	}
	path = node[to].path
	rate = node[to].rate
	return
}

func isVisited(path []int, node int) bool {
	for _, v := range path {
		if v == node {
			return true
		}
	}
	return false
}
