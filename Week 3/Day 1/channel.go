package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var pesan = make(chan string)

	var sayHello = func(who string) {
		var data = fmt.Sprintf("hello %s", who)

		pesan <- data
	}

	go sayHello("Nobita")
	go sayHello("Suneo")
	go sayHello("Giant")
	go sayHello("Sisuka")
	go sayHello("doraemon")

	var pesan1 = <-pesan
	fmt.Println(pesan1)

	var pesan2 = <-pesan
	fmt.Println(pesan2)

	var pesan3 = <-pesan
	fmt.Println(pesan3)

	// var pesan4 = <-pesan
	// fmt.Println(pesan4)

	// var pesan5 = <-pesan
	// fmt.Println(pesan5)
}
