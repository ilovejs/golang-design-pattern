package _map

import (
	"strconv"
	"sync"
	"testing"
)

const (
	NumOfReader = 100
	NumOfWriter = 10
)

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < NumOfWriter; j++ {
			wg.Add(1)
			go func() {
				for k := 0; k < 100; k++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for l := 0; l < NumOfReader; l++ {
			wg.Add(1)
			go func() {
				for m := 0; m < 100; m++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	b.Run("Map with RWLock", func(b *testing.B) {
		hm := CreateRwLockMap()
		benchmarkMap(b, hm)
	})

	b.Run("sync.map", func(b *testing.B) {
		hm := CreateSyncMapBenchmarkAdapter()
		benchmarkMap(b, hm)
	})

	b.Run("concurrent map", func(b *testing.B) {
		superman := CreateConcurrentMapBenchmarkAdapter(199)
		benchmarkMap(b, superman)
	})
}
