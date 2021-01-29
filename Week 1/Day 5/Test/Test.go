package main

import (
	"fmt"
	"regexp"
)


type Phone struct{
	name string
	email string
	number string
}


// func tambahdataSementara(name,email string, number int64) (phoneBookSementara string){
// 	phoneBookSementara =  name + email + fmt.Sprintf("%f", number) 
// 	return
// }

// func tambahData(number int64,name,email string) (phoneBookSementara string){
// 	phoneBookSementara = tambahdataSementara(name,email,number)
// 	return
// }


func menu(){
	fmt.Println("-----------------------------------------")
	fmt.Println("|Selamat Datang Di Simple App PhoneBook  |")
	fmt.Println("|[1].InputData                           |")
	fmt.Println("|[2].ListBook                            |")
	fmt.Println("|[3].UpdateEmail                         |")
	fmt.Println("|[4].UpdateNomor                         |")
	fmt.Println("|[!]. untuk keluar                       |")
	fmt.Println("-----------------------------------------")
	fmt.Print("Masukan Menu pilihan Anda :")
}
func checkHuruf(name string) bool {
	isStringAlphabetic,_ := regexp.MatchString(`^[a-z A-Z]+$`, name)
	return isStringAlphabetic
}

func checkNomor(nomor string) bool {
	isStringNumeric,_ := regexp.MatchString(`^[\d]+$`, nomor)
	return isStringNumeric
}

func checkEmail(email string) bool{
	isStringValid,_ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",email)
	return isStringValid
}


func main() {
	var inputMenu,inputDalamMenu,inputNamaSementara,InputEmailSementara string
	var inputNomorSementara string
	var ListBook = []Phone{}
	var phone Phone
	for {
		menu()
		fmt.Scanln(&inputMenu)
		if inputMenu == "!"{
			fmt.Println("Anda Sudah Keluar Dari Apilikasi")
			break
		}else if inputMenu == "1" {
			fmt.Println()
			fmt.Println("======================================")
			fmt.Println("|Anda masuk Ke Menu Tambah Data  Nomor|")
			fmt.Println("======================================")
			fmt.Println()
			for {
				fmt.Println()
				fmt.Println("--------------------")
				fmt.Print("Masukan Nama Anda : ")
				fmt.Scanln(&phone.name)
				if  len(phone.name) <=5 && checkHuruf(phone.name){
					fmt.Println("[Minimal 5 Huruf Dan Tidak Boleh Angka !]")
					break
				}else {
					fmt.Println("Nama Anda adalah:", phone.name)
				}
				fmt.Println()
				fmt.Println("-------------------")
				fmt.Print("Masukan Email Anda : ")
				fmt.Scanln(&phone.email)
				if checkEmail(phone.email){
					fmt.Println("Email Anda adalah:", phone.email)
				}else{
					fmt.Println("[Email Yang Anda Masukan Tidak Valid!]")
					break
				}
				fmt.Println()
				fmt.Println("-------------------")
				fmt.Print("Masukan Nomer Anda :  ")
				fmt.Scanln(&phone.number)
				if len(phone.number) <= 10 && checkNomor(phone.number) == false{
					fmt.Println("[Nomor Tidak Boleh Kurang Dari 10 Digit dan Tidak Boleh Huruf!!]")
					fmt.Println("Nomor Anda Adalah" , phone.number)
					fmt.Println()
					break
				}
				fmt.Println()
				ListBook= append(ListBook, phone)
				fmt.Println("===================================================================")
				fmt.Println("Nama :",phone.name,"Email :", phone.email, "Nomor : ", phone.number)
				fmt.Println("===================================================================")
				fmt.Println("Nomor Sudah Tersimpan, Ketik [!] Untuk Kembali Ke Menu")
				fmt.Println("===================================================================")
				fmt.Scanln(&inputDalamMenu)
				if inputDalamMenu == "!"{
					fmt.Println()
					fmt.Println("=======================================================================")
					fmt.Println("|                      Anda Keluar dari Menu                          |")
					fmt.Println("=======================================================================")
					fmt.Println()

					break
				}else {
					fmt.Println("")
					continue
				}
			}
		} else if inputMenu == "2" {
			fmt.Println()
			fmt.Println("===================================================================")
			fmt.Println("|                        Anda Masuk Ke Menu ListBook               |")
			fmt.Println("===================================================================")
			fmt.Println(ListBook)
			fmt.Println("Apakah Anda Ingin Keluar Dari Menu ? Ketik [!] Untuk Kembali Ke Menu Ketik [Y] Untuk Keluar Apilikasi")
			fmt.Scanln(&inputDalamMenu)
				if inputDalamMenu == "!"{
					fmt.Println()
					fmt.Println("=======================================================================")
					fmt.Println("|                  Anda Keluar Dari Menu                              |")
					fmt.Println("=======================================================================")
					menu()
				}else if inputDalamMenu == "Y" {
					fmt.Println()
					fmt.Println("=======================================================================")
					fmt.Println("|                  Anda Keluar Dari Apilikasi                          |")
					fmt.Println("=======================================================================")
					break
				}
		}else if inputMenu == "3"{
			fmt.Println()
			fmt.Println("===================================================================")
			fmt.Println("|                Anda Masuk Ke menu Update Email                  |")
			fmt.Println("===================================================================")
			fmt.Print("Masukan Nama List Contact Yang Akan Di Update :   ")
			fmt.Scanln(&inputNamaSementara)
			fmt.Println()
			fmt.Print("Masukan Email Baru :")
			fmt.Scanln(&InputEmailSementara)
			if !checkEmail(InputEmailSementara) {
				fmt.Println("[Email Yang Anda Masukan Tidak Valid!]")
				break
			}
			for i:= 0; i < len(ListBook);i++{
				if ListBook[i].name == inputNamaSementara{
					ListBook[i].email = InputEmailSementara
					fmt.Println("===============================================")
					fmt.Println("|             Email Sudah TerUpdate           |")
					fmt.Println("==============================================")
				}else{
					fmt.Println("Error")
				}
				break			
			}
		}else if inputMenu == "4"{
			fmt.Println()
			fmt.Println("===================================================================")
			fmt.Println("|                  Anda Masuk Ke Menu Update Nomor                 |")
			fmt.Println("===================================================================")
			fmt.Print("Masukan Nama List Contact Yang Akan Di Update : ")
			fmt.Scanln(&inputNamaSementara)
			fmt.Println()
			fmt.Print("Masukan nomor baru :  ")
			fmt.Scanln(&inputNomorSementara)
			if !checkNomor(inputNomorSementara){
				fmt.Println("[Masukan Minimal 10 Digit dan Tidak Boleh Huruf!]")
			}

			for i:= 0; i < len(ListBook);i++{
				if ListBook[i].name == inputNamaSementara{
					ListBook[i].number = inputNomorSementara
					fmt.Println("===============================================")
					fmt.Println("|             Nomor Sudah TerUpdate           |")
					fmt.Println("==============================================")
				}else{
					fmt.Println("Error")
				}
				break			
			}
		}
	}
}