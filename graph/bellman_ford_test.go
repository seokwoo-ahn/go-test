package graph

import (
	"fmt"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	edge := [][]float64{
		{-1, 1, -1, -1},
		{1.2, -1, 0.9, 1.2},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
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

// loop를 돌았을때 swap pool의 ratio가 1이 될때까지만 돌면 돈이 계속 증가함
func TestBellmanFordLoop(t *testing.T) {
	pool1TokenA := float64(10000000)
	pool1TokenB := float64(12000000)

	pool2TokenB := float64(10000000)
	pool2TokenA := float64(12000000)

	amount := float64(100000)
	resA := amount
	resB := float64(0)

	for i := 0; i < 10; i++ {
		ratio1 := float64(pool1TokenB) / (float64(pool1TokenA) + float64(resA))
		pool1TokenA += resA
		pool1TokenB -= resA * ratio1

		resB += resA * ratio1
		resA -= resA

		ratio2 := float64(pool2TokenA) / (float64(pool2TokenB) + float64(resB))
		pool2TokenB += resB
		pool2TokenA -= resB * ratio2

		resA += resB * ratio2
		resB -= resB

		fmt.Println(pool1TokenA, pool2TokenB)
		fmt.Println("resA", resA)
	}
}

// swap 비율은 1:1로 수렴함
func TestSwapLoopWithFixedAmount(t *testing.T) {
	pool1TokenA := float64(11500000)
	pool1TokenB := float64(12000000)

	amount := float64(10000)
	resA := amount
	resB := float64(0)

	for i := 0; i < 20000; i++ {
		ratio1 := float64(pool1TokenB) / (float64(pool1TokenA) + float64(amount))
		resB += amount * ratio1
		resA -= amount
		pool1TokenA += amount
		pool1TokenB -= amount * ratio1

		fmt.Println(pool1TokenA, pool1TokenB)
		fmt.Println("ratio1:", ratio1)
		fmt.Println("resA", resA)

		ratio2 := float64(pool1TokenA) / (float64(pool1TokenB) + float64(amount))
		resA += amount * ratio2
		resB -= amount
		pool1TokenB += amount
		pool1TokenA -= amount * ratio2

		fmt.Println(pool1TokenA, pool1TokenB)
		fmt.Println("ratio2:", ratio2)
		fmt.Println("resB", resB)

		fmt.Println("total:", resA+resB+float64(pool1TokenA)+float64(pool1TokenB))
		fmt.Println("client res:", resA+resB)
	}
}

// client의 돈은 그대로 유지
func TestSwapLoop(t *testing.T) {
	pool1TokenA := float64(11500000)
	pool1TokenB := float64(12000000)

	amount := float64(10000)
	resA := amount
	resB := float64(0)

	for i := 0; i < 10; i++ {
		ratio1 := float64(pool1TokenB) / (float64(pool1TokenA) + float64(resA))
		pool1TokenA += resA
		pool1TokenB -= resA * ratio1
		resB += resA * ratio1
		resA -= resA

		fmt.Println(pool1TokenA, pool1TokenB)
		fmt.Println("ratio1:", ratio1)
		fmt.Println("resA", resA)

		ratio2 := float64(pool1TokenA) / (float64(pool1TokenB) + float64(resB))
		pool1TokenB += resB
		pool1TokenA -= resB * ratio2
		resA += resB * ratio2
		resB -= resB

		fmt.Println(pool1TokenA, pool1TokenB)
		fmt.Println("ratio2:", ratio2)
		fmt.Println("resB", resB)

		fmt.Println("total:", resA+resB+float64(pool1TokenA)+float64(pool1TokenB))
		fmt.Println("client res:", resA+resB)
	}
}
