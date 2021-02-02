package main

import (
	"fmt"
	"os"
	_runtime "runtime/pprof"
)

var (
	cpuProfile = "cpu.prof"
)

func main() {
	cpuPros, err := os.Create(cpuProfile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cpuPros.Close()

	_runtime.StartCPUProfile(cpuPros)
	defer _runtime.StopCPUProfile()
	loopData()
}

func loopData() {
	var slice []int
	for i := 0; i < 100000; i++ {
		slice = append(slice, i)
	}
}

func processBigArray() {
	for i := 0; i < 1000000; i++ {
		arr := bigArray()
		if arr == nil {
			fmt.Println("Array is nil")
		}
	}
}

func bigArray() *[]int {
	s := make([]int, 1000000)
	return &s
}
