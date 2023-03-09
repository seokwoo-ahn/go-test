package algorithm

import (
	"fmt"
	"testing"
)

func TestRabinKarp(t *testing.T) {
	text := "abcabcbabaababc"
	pattern := "babc"

	fmt.Println(RabinKarp(text, pattern))
}
