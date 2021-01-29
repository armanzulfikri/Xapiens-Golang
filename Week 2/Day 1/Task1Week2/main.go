package main

import (
    "fmt"
    "math/rand"
)

type LuckyNumber struct {
    number int
}

type LuckyNumberKumpulan struct{
	listNumberKumpulan []LuckyNumber
}

func main() {
	var nilai []LuckyNumber
    for i := 0; i <= 200; i++ {
        nilai = append(nilai, LuckyNumber{rand.Intn(100)})
    }

	kelompok1 := nilai[0:39]
	kelompok2 := nilai[40:79]
	kelompok3 := nilai[80:129]
	kelompok4 := nilai[130:159]
	kelompok5 := nilai[160:200]
	
	
	fmt.Println("nilai Kelompok 1 :")
	fmt.Print(kelompok1)
	
	fmt.Println("nilai Kelompok 2 :")
	fmt.Println(kelompok2)

	
	fmt.Println("nilai Kelompok 3 :")
	fmt.Println(kelompok3)

	fmt.Println("nilai Kelompok 4 :")
	fmt.Println(kelompok4)
	
	fmt.Println("nilai Kelompok 5 :")
	fmt.Println(kelompok5)


	

}


