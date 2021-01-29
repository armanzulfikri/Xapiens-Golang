package main
import "fmt"

func main() {
	suhu := 20

	if suhu <= 0 {
		fmt.Println("beku")
	}else if suhu > 0  && suhu <= 100 {
		fmt.Println("cair")
	}else{
		fmt.Println("Lentur")
	}


	nama := "arman"

	if length := len(nama); length <=5 {
		fmt.Println(length, "nama yang di input terlalu pendek")
	}else  {
		fmt.Println (length, "nama yang di input sudah benar ")
	}

	switch nama {
	case "budi":
		fmt.Println("Hallo Budi")
	case "dwi" : 
		fmt.Println("Halo dwi")
	default :
		fmt.Println("nama belem terdaftar")
	}


	switch  length := len(nama); length <= 5 {
	case true:
		fmt.Println("nama terlalu pendek")
	default :
		fmt.Println ("nama sudah benar")
	}

	nama = "cahyo Purnomo"
	suhu = -3
	length := len(nama)

	switch  {
	case length <= 5:
		fmt.Println("Nama terlalu pendek")
	case length >5 :
		fmt.Println("nama sudah benar")
	case suhu <=0 : 
		fmt.Println("Beku")
	}






	///for 
	count := 1// kondisi yang belum diketahui
	for count <=5 {
		fmt.Print("Perulangan ke", count)
		count ++
	}

	for i:=1; i<=5; i++{
		fmt.Println("perulangan ke ", i)
	}

	slice := []string{"solo","klaten","sleman"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "value", value)
	}



	nilai:=1 //kondisi sudah diketahui
	for {
		
		if nilai == 3 {
			nilai++
			continue
		}
		fmt.Println("perulangan ke", nilai)
		nilai++

		if nilai ==10 {
			break
		}
		
	}
}