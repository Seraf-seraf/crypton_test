// Кравчук Серафим Сергеевич
package main

import (
	"math/rand"
	"sync"
	"testing"
)

func TestProtectedMap(t *testing.T) {
	pm := ProtectedMap{
		m: make(map[int]int),
	}
	wg := &sync.WaitGroup{}

	// 1776. Год принятия Декларации независимости США
	const year = 1776
	jobs := make(chan int, year)

	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for key := range jobs {
				pm.Touch(key)
			}
		}()
	}

	for i := 1; i <= 3; i++ {
		seen := make(map[int]struct{})

		for len(seen) < year {
			num := rand.Intn(year) + 1

			if _, ok := seen[num]; !ok {
				seen[num] = struct{}{}
				jobs <- num
			}
		}
	}

	close(jobs)
	wg.Wait()

	if pm.keyAccessCounter != year*3 {
		t.Errorf("keyAccessCounter error")
	}

	if pm.newKeysCounter != year {
		t.Errorf("newKeysCounter error")
	}

	for i := 1; i <= year; i++ {
		if pm.m[i] != 3 {
			t.Errorf("m[k] = %d", i)
		}
	}
}
