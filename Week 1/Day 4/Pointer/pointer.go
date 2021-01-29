package main

import "fmt"

type Address struct {
		city,province,country string
}

func main() {

	address1 := Address{"Kalten","jateng","indonesia"}
	address2 := address1

	address1.city = "bandung"

	fmt.Println(address1)
	fmt.Println(address2)


	var address3 *Address = &address1
	//addrres3 pasti akan menerima perubahan apapaun dari address1 dan sebaliknya
	
	address3.city = "solo"

	fmt.Println("1 1",address1)
	fmt.Println("1 2",address2)
	fmt.Println("1 3",address3)


	var address4 = &Address{"balikpapan","kaltim","indonesia"}
	fmt.Println("2 1",address1)
	fmt.Println("2 2",address2)
	fmt.Println("2 3",address3)
	fmt.Println("2 4",address4)

	address3 = &Address{"banjar","kalsel","indonesia"}
	fmt.Println("3 1",address1)
	fmt.Println("3 2",address2)
	fmt.Println("3 3",address3)
	fmt.Println("3 4",address4)


	address5 := &address1
	*address5 = Address{"surabaya","jatim","indonesia"}
	fmt.Println("4 1",address1)
	fmt.Println("4 2",address2)
	fmt.Println("4 3",address3)
	fmt.Println("4 4",address4)
	fmt.Println("4 5", address5)

	// address6 := &address1
	// *address6 = Address{"jogja","DIY","indonesia"}
	// // address6.city = "Jogja"
	// fmt.Println("5 1",address1)
	// fmt.Println("5 2",address2)
	// fmt.Println("5 3",address3)
	// fmt.Println("5 4",address4)
	// fmt.Println("5 5", address5)
	// fmt.Println("5 6", address6)

		//membuat address yang kosong
	// address7 *Address = new(Address)
	// fmt.Prinln("5 7", address7)
	


	changeCity(&address2)
	changeCity(&address1) /// 1 3 5 berubah karena mereka merujuk ke pointer 1

	fmt.Println("6 1",address1)
	fmt.Println("6 2",address2)
	fmt.Println("6 3",address3)
	fmt.Println("6 4",address4)
	fmt.Println("6 5", address5)


	address2.city = "kudus"
	fmt.Println("7 1",address1)
	fmt.Println("7 2",address2)
	fmt.Println("7 3",address3)
	fmt.Println("7 4",address4)
	fmt.Println("7 5", address5)

	
	cahyo := Person{"Cahyo"}
	cahyo.school()
	fmt.Println(cahyo.name)

}
func changeCity(address *Address){
	address.city = "semarang"
}

type Person struct{
	name string
}
func (person *Person)school(){
	person.name = "Arman"
}