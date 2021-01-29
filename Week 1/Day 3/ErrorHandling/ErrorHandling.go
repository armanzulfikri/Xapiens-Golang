package main

import (
	"fmt"
	"errors"
)

func logrecover(){
	fmt.Println("selesai memanggil fungsi")
	pesanError := recover()
	if pesanError != nil {
		fmt.Println("Terjadi error", pesanError)
	}
}


func pembagian(value int){
	defer logrecover()
	fmt.Println("ini pembagian")
	total := 10/value
	fmt.Println(total)
}


func pembagianDenganError(value int) (int, error){

	if value == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	}else{
	return (10/value), nil
	}
}

func main() {
//pemnbagian 0
//tipe data tidak sesuai
//range tipe data tidak seusai
//defer --> try
//panic --> error
//recover --> cacth
pembagian(0)
hasil, error := pembagianDenganError(9)
fmt.Println("hasil", hasil, "","eror",error)


}