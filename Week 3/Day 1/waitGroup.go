package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print(wg *sync.WaitGroup, pesan string) {
	defer wg.Done()
	fmt.Println(pesan)
}
func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)
		wg.Add(1)
		go print(&wg, data)
	}
	wg.Wait()
}
