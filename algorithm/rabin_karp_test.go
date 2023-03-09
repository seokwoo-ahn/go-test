package algorithm

import (
	"fmt"
	"testing"
)

func TestRabinKarp(t *testing.T) {
	text := "abcabcbabaababc"
	pattern := "cba"

	fmt.Println(RabinKarp(text, pattern))
}
