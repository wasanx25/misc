package main

import (
	"fmt"
	"runtime"
)

var mem runtime.MemStats

func PrintMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Printf(
		"Alloc: %d KB, TotalAlloc: %d KB, HeapAlloc: %d KB, HeapSys: %d KB\n",
		mem.Alloc/1024,
		mem.TotalAlloc/1024,
		mem.HeapAlloc/1024,
		mem.HeapSys/1024,
	)
}

func main() {
	PrintMemory()
	size := 100000000
	fmt.Printf("slice size : %d \n", size)
	b := make([]int, size)
	for i, _ := range b {
		b[i] = i
	}
	PrintMemory()
}
