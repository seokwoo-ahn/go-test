package algorithm

import (
	"fmt"
	"testing"
)

func TestPrefixSuffixTable(t *testing.T) {
	str := "abcabcbabaababc"
	fmt.Println(FillKmpTable(str))
}

func TestKMP(t *testing.T) {
	str := "abcabcbabaababc"
	pattern := "cba"

	fmt.Println(KMP(str, pattern))
}
