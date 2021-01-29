package main

import "fmt"

func main(){
	fmt.Println("Tipe Data")
	fmt.Println("Bilangan Bulat", 6)
	fmt.Println ("Bilangan Desimal", 6.5)
	fmt.Println("Bilangan Boolean", true)
	fmt.Println("Ini Tipe Data String")


	var nama string
	var ipk float32

	nama = "Arman Zulfikri Fardiansyah"
	ipk = 3.99

	tinggi := 181 //deklarasi tanpa medeklare tipe datanya

	total := ipk * 3;

	fmt.Println("Nama Saya adalah", nama)
	fmt.Println("Total IPK", total)
	fmt.Println("tinggi", tinggi,"cm")

	text := "ini string"
	fmt.Println(text[1])
	fmt.Println(string(text[1]))


	var angka int16
	var angkaConvert int8

}