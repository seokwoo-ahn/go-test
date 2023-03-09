package hack

import (
	"fmt"
	"math"
	"math/big"
)

func Decode(q *big.Int, n, key2 int64) *big.Int {
	mul := new(big.Int).Mul(q, big.NewInt(n+1))

	for i := 1; i < 95; i++ {
		key2Pow := math.Pow(float64(key2), float64(n+1))
		key1Pow := math.Pow(float64(i), float64(n+1))
		sub := key1Pow - key2Pow
		bigSub := big.NewInt(int64(sub))
		mod := new(big.Int).Mod(mul, bigSub)

		if mod.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("Key1:", i)
			return new(big.Int).Div(mul, big.NewInt(int64(sub)))
		}
	}
	return big.NewInt(0)
}
