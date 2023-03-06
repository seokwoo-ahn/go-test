package algorithm

import (
	"fmt"
	"testing"
)

func TestPrefixSuffixTable(t *testing.T) {
	str := "abcabcbabaababc"
	fmt.Println(*FillKmpTable(str))
}
