package main

import (
	"fmt" 
	"math/rand"
	// "time"
	"sort"
)

func main() {
	min := 20
	maxNilai:= 100
	max := 100

	 
	nilai :=[]int{}
	for i:=1;i<=max;i++{ //membuat 100 random nilai
	nilai = append(nilai,rand.Intn(maxNilai-min)+min)
	}
	fmt.Println(nilai)

///
	kelompok1 := nilai[0:33]//slicing nilai 
	kelompok2 := nilai[34:67]
	kelompok3 := nilai[68:100]


	rata1 := 0 
	total1 := 0
	minimum1 :=0
	maksimum1:=0
	
	for i:=0;i<len(kelompok1);i++{ //mencari total kelompok 1
		total1 += kelompok1[i]
		rata1= total1/len(kelompok1)
	}
	for i, e:= range kelompok1{ //mencari minimum
		if i==0 || e<minimum1{
			minimum1 = e
		}
	}
	for i, e:= range kelompok1{ //mencari maksimum
		if i==0 || e>maksimum1{
			maksimum1 = e
		}
	}
	fmt.Println("Kelompok 1")
	fmt.Println("total:", total1)
	fmt.Println("rata2:",rata1)
	fmt.Println("min:",minimum1)
	fmt.Println("max:",maksimum1)
	sort.Ints(kelompok1)
	fmt.Println("urut :",kelompok1)


	rata2 := 0 
	total2 := 0
	minimum2 :=0
	maksimum2:=0
	
	for i:=0;i<len(kelompok2);i++{ //mencari total kelompok 2
		total2 += kelompok2[i]
		rata2= total2/len(kelompok2)
	}
	for i, e:= range kelompok2{ //mencari minimum
		if i==0 || e<minimum2{
			minimum2 = e
		}
	}
	for i, e:= range kelompok2{ //mencari maksimum
		if i==0 || e>maksimum2{
			maksimum1 = e
		}
	}
	fmt.Println("Kelompok 2")
	fmt.Println("total:", total2)
	fmt.Println("rata2:",rata2)
	fmt.Println("min:",minimum2)
	fmt.Println("max:",maksimum2)
	sort.Ints(kelompok2)
	fmt.Println("urut :",kelompok2)



	rata3 := 0 
	total3 := 0
	minimum3 :=0
	maksimum3:=0
	
	for i:=0;i<len(kelompok3);i++{ //mencari total kelompok 3
		total3 += kelompok3[i]
		rata3= total3/len(kelompok3)
	}
	for i, e:= range kelompok3{ //mencari minimum
		if i==0 || e<minimum3{
			minimum3 = e
		}
	}
	for i, e:= range kelompok3{ //mencari maksimum
		if i==0 || e>maksimum3{
			maksimum3 = e
		}
	}
	fmt.Println("Kelompok 3")
	fmt.Println("total:", total3)
	fmt.Println("rata2:",rata3)
	fmt.Println("min:",minimum3)
	fmt.Println("max:",maksimum3)
	sort.Ints(kelompok3)
	fmt.Println("urut :",kelompok3)




	




}