package crypto

import (
	"fmt"
	"testing"
)

func TestCryptoRandom(t *testing.T) {
	try := 0
	for try < 10 {
		// 실행할때마다 다른 값
		fmt.Println("CryptoRandom:", CryptoRandom())
		// 실행할때마다 같은 값
		fmt.Println("Normal Random:", Random())
		try++
	}
}
