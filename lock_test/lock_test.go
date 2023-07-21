package lock_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var data int = 0
	var rwMutex = new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			time.Sleep(3 * time.Second)
			fmt.Println("read 1 : ", data)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2 : ", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1
			fmt.Println("write  : ", data)
			time.Sleep(1 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
