package exporter

import (
	"fmt"
	"runtime"
)

func memoryUsage() {
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)

	fmt.Printf("Memoria usada: %.2f Mb\n", float64(memory.Alloc)/1024/1024)
	fmt.Printf("Memoria Total del programa: %.2f Mb\n", float64(memory.Sys)/1024/1024)
	fmt.Printf("Heap en uso: %.2f Mb\n", float64(memory.HeapInuse)/1024/1024)
	fmt.Println("")
}
