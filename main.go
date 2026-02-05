// Кравчук Серафим Сергеевич
package main

import (
	"fmt"
	"sync"
)

type ProtectedMap struct {
	mu sync.Mutex
	m  map[int]int

	keyAccessCounter int
	newKeysCounter   int
}

func (pm *ProtectedMap) Touch(v int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, ok := pm.m[v]; !ok {
		pm.m[v] = 0
		pm.newKeysCounter++
	}

	pm.m[v]++
	pm.keyAccessCounter++
}

func main() {
	pm := &ProtectedMap{
		m: make(map[int]int),
	}

	// 1776. Год принятия Декларации независимости США
	pm.Touch(1776)

	fmt.Println("pm.keyAccessCounter:", pm.keyAccessCounter)
	fmt.Println("pm.newKeysCounter:", pm.newKeysCounter)
}
