package hack

import (
	"fmt"
	"math/big"
	"testing"
)

func TestDecode(t *testing.T) {
	test := new(big.Int)
	testString := "-200640142664324295933714"
	a, _ := test.SetString(testString, 10)
	res := Decode(a, 3, 95)
	fmt.Println(res)
}
