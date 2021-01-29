package main

import "fmt"
 
func main(){
	var cities[3] string
	cities[0]="bali"
	cities[1]="sleman"
	cities[2]="surabaya"

	fmt.Println(cities)
	fmt.Println(cities[0])

	var ages = [3]int{
		20,21,25,
	}

	fmt.Println(ages)
	fmt.Println(ages[0])
	fmt.Println(len(ages))

	var angkaArray = [...]string{
		"satu",
		"dua",
		"tiga",
	}

	fmt.Println(angkaArray)
	fmt.Println(len(angkaArray))
	





	//slice
	var angkaSlice = []string{
		"satu",
		"dua",
		"tiga",
	}


	fmt.Println(angkaSlice)
	fmt.Println(len(angkaSlice))
	fmt.Println(cap(angkaSlice))



	numberSlice := []int{
		1,2,3,4,5,
	}

	numberSlice = append(numberSlice, 6)
	fmt.Println(numberSlice)
	fmt.Println(len(numberSlice))
	fmt.Println(cap(numberSlice))




	//slice dari array 
	fmt.Println("---------------------------")
	fmt.Println("Ini adalah Slice dari Array")
	var angkaBiasa = [...]string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
	}


	var slice =  angkaBiasa[1:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))




	newSlice := make([]string,3,6)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))





	//Slice reference to array
	fmt.Println("Slice reference to array")
	var alphabet = [...]string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		
	}

	var alphabetSlice = alphabet[0:5]
	fmt.Println(alphabet)
	fmt.Println(alphabetSlice)
	
	alphabetSlice[1]= "b2"

	fmt.Println(alphabet)
	fmt.Println(alphabetSlice)



	var alphabetSlice2 = alphabet[0:5]
	fmt.Println(alphabet)
	fmt.Println(alphabetSlice)
	fmt.Println(alphabetSlice2)

	alphabetSlice[1] = "b3"
	fmt.Println(alphabet)
	fmt.Println(alphabetSlice)
	fmt.Println(alphabetSlice2)

	fmt.Println(len(alphabetSlice),cap(alphabetSlice))
	alphabetSlice =  append(alphabetSlice,"f2")
	fmt.Println(alphabet)
	fmt.Println(alphabetSlice)
	fmt.Println(alphabetSlice2)
	alphabetSlice2 = append(alphabetSlice2, "f3")
	fmt.Println(alphabet)
	fmt.Println(alphabetSlice)
	fmt.Println(alphabetSlice2)
	alphabetSlice =append(alphabetSlice,"g4")
	fmt.Println(alphabet)
	fmt.Println(alphabetSlice)
	fmt.Println(alphabetSlice2)





	//map
	fmt.Println("==============")
	fmt.Println("Map")

	user := map[string]string{
		"nama" : "arman",
		"kota" : "banyuwangi",
	}

	fmt.Println(user)
	fmt.Println(len(user)) 

	delete(user, "kota")


	fmt.Println(user)
	fmt.Println(len(user))

	var karyawan = make(map[string]string)
	karyawan["NIP"]= "12345"

	//map with slice

	var mahasiswa =[]map[string]string{
		map[string]string{"nama":"arman","nim":"1234"},
		map[string]string{"nama":"agus","nim":"1234"},
		
	}
	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa[0]["nama"])


}