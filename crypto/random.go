package crypto

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func CryptoRandom() int {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		fmt.Println("fail to load crypto random value")
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
	return rand.Int()
}

func Random() int {
	rand.Seed(int64(64))
	return rand.Int()
}
