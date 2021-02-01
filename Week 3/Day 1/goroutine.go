package main

import (
	"fmt"
	"runtime"
)

func showData(maksIndex int, message string) {
	for i := 0; i < maksIndex; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	go showData(100, "Tes Goroutine")
	showData(100, "tanpa Goroutine")

	var input string
	fmt.Scanln(&input)
}
