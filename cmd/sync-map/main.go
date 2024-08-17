package main

import (
	"fmt"
	"sync"
	"time"

	"go-playground/pkg/safemap"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Map
	sm := safemap.New[int, string]()

	// SYNC MAP
	nowForSyncMap := time.Now()

	fmt.Println("storing data in the SYNC MAP")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go storeSyncMapData(&wg, &m, i)
	}
	wg.Wait()

	fmt.Println("printing data from 'for' loop")
	for i := 0; i < 5; i++ {
		printSyncMapData(&m, i)
	}

	fmt.Println("printing data from 'map.Range'")
	m.Range(syncMapPrinter) // Equivalent to 'for k, v := range m {...}' of normal maps

	fmt.Println("Time with SYNC MAP: ", time.Since(nowForSyncMap))

	// SAFE MAP
	nowForSafeMap := time.Now()

	fmt.Println("storing data in the SAFE MAP")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go storeSafemapData(&wg, sm, i)
	}
	wg.Wait()

	fmt.Println("printing data from 'for' loop")
	for i := 0; i < 5; i++ {
		printSafemapData(sm, i)
	}

	fmt.Println("printing data from 'map.Range'")
	sm.Range(safeMapPrinter) // Equivalent to 'for k, v := range m {...}' of normal maps

	fmt.Println("Time with SAFE MAP: ", time.Since(nowForSafeMap))
}

func storeSyncMapData(wg *sync.WaitGroup, m *sync.Map, key int) {
	defer wg.Done()
	m.Store(key, fmt.Sprintf("test %v", key))
}

func printSyncMapData(m *sync.Map, key int) {
	v, ok := m.Load(key)
	if !ok {
		fmt.Println("error loading key", key)
	}
	fmt.Println(v)
}

func syncMapPrinter(k any, v any) bool {
	fmt.Println(v)
	return true
}

func storeSafemapData(wg *sync.WaitGroup, m *safemap.SafeMap[int, string], key int) {
	defer wg.Done()
	m.Store(key, fmt.Sprintf("test %v", key))
}

func printSafemapData(m *safemap.SafeMap[int, string], key int) {
	v, ok := m.Load(key)
	if !ok {
		fmt.Println("error loading key", key)
	}
	fmt.Println(v)
}

func safeMapPrinter(k int, v string) {
	fmt.Println(v)
}
