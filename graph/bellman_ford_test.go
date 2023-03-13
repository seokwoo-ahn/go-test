package graph

import (
	"fmt"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	edge := [][]float64{
		{-1, 1, -1, -1},
		{-1, -1, 0.9, 1.2},
		{1.3, -1, -1, -1},
		{-1, -1, 1, -1},
	}
	from := 0
	to := 3

	rate, path, err := BellmanFord(edge, from, to)
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println(path)
		return
	}
	fmt.Println(rate)
	fmt.Println(path)
}
