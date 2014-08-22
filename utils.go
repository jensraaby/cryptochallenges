// Package  provides ...
package main

import "sort"

const TopLetters string = "ETAOIN SHRDLU"

// SortedByteMap is used for emitting bytes sorted descending by count
type SortedByteMap struct {
	m map[byte]int
	c []byte
}

func (sm *SortedByteMap) Len() int {
	return len(sm.m)
}
func (sm *SortedByteMap) Less(i, j int) bool {
	return sm.m[sm.c[i]] > sm.m[sm.c[j]]
}
func (sm *SortedByteMap) Swap(i, j int) {

	sm.c[i], sm.c[j] = sm.c[j], sm.c[i]
}
func SortedKeys(m map[byte]int) []byte {
	sm := new(SortedByteMap)
	sm.m = m
	sm.c = make([]byte, len(m))
	i := 0
	for k, _ := range m {
		sm.c[i] = k
		i++
	}
	sort.Sort(sm)
	return sm.c
}
