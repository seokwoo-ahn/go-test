package map_test

import (
	"math/rand"
	"testing"
)

func TestThreadSafe(t *testing.T) {
	t.Run("thread safe", func(t *testing.T) {
		mw := NewMapWithMutex()
		for i := 0; i < 100; i++ {
			go mw.Set(i, rand.Int())
			go mw.Get(0)
		}
	})

	t.Run("thread unsafe", func(t *testing.T) {
		m := NewMapWithOutMutex()
		for i := 0; i < 100; i++ {
			go m.Set(i, rand.Int())
			go m.Get(0)
		}
	})
}
