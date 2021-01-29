package main

import (
	"fmt"
	"math/rand"
	"sort"
)


func min(kelompok ...int) int {
	minimum := kelompok[0]
	for _, e := range kelompok {
		if e < minimum {
			minimum = e
		}
	}
	return minimum
}

func max(kelompok ...int) int {
	maximum := kelompok[0]
	for _, e := range kelompok {
		if e > maximum{
			maximum = e
		}
	}
	return maximum
}

func sum(kelompok ...int)int {
	total := kelompok[0]
	for _, e := range kelompok {
		total += e
	}
	return total
}


func average (total int, kelompok ...int)float32 {
	return float32(total) / float32(len(kelompok))
}











func report(typeKelompok int, kelompok ...int){
	total := sum(kelompok...)
	rata := average(total,kelompok...)
	min := min(kelompok...)
	max := max (kelompok...)

	fmt.Println("Kelompok ", typeKelompok)
	fmt.Println("total:", total)
	fmt.Println("rata2:",rata)
	fmt.Println("min:",min)
	fmt.Println("max:",max)
	sort.Ints(kelompok)
	fmt.Println("urut :",kelompok)

}

func sliceKelompok(arr ...int)(kelompok1,kelompok2,kelompok3 []int){
	for i, e:= range arr{
		if i%3 == 0 {
			kelompok1 = append(kelompok1, e)
		}else if i%3 == 1{
			kelompok2 = append(kelompok2, e)
		}else if  i%3 == 2 {
			kelompok3 = append(kelompok3, e)
		}

	}
	return

}

func inisialKelompok(min int, max int, maxNilai int)(nilai []int){
	for i := 1; i <= max; i++ {
		nilai = append(nilai, rand.Intn(maxNilai-min)+min)
	}
	fmt.Println(nilai)
	return

}

func main() {
	arr := inisialKelompok(20, 100, 100)
	kelompok1,kelompok2,kelompok3 := sliceKelompok(arr...)
	report(1, kelompok1...)
	report(2, kelompok2...)
	report(3, kelompok3...)
}


