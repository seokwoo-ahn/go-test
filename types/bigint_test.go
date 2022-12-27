package types

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBigIntMul(t *testing.T) {
	testInt := int64(9223372036854775807)
	fmt.Println(testInt)
	bigInt1 := big.NewInt(testInt)
	fmt.Println(bigInt1)
	bigInt2 := bigInt1.Mul(bigInt1, bigInt1)
	fmt.Println(bigInt2) // 다른 언어에서의 big int range 초과 가능
	fmt.Println(bigInt1) // bigInt1 값 변경
}
