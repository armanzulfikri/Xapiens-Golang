### Belajar Golang Task 1 Week 1

### Number 1
-Golang (atau biasa disebut Go) adalah bahasa pemrograman baru yang dikembangkan di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson pada tahun 2007 dan mulai di perkenalkan ke publik tahun 2009. Penciptaan bahasa Go didasari bahasa C dan C++ oleh karena itu gaya sintaknya mirip.

-Kelebihan Go :
1.Mendukung konkurensi di level bahasa dengan mengaplikasian cukup mudah
2.Memiliki garbage collerctor
3.Proses kompilasi sangat cepat
4.Bukan bahasa pemrograman yang hirarkial, menjadikan developer tidak perlu ribet lagi memikirkan segmen OOP-nya
5.Package/modul yang disediakan terbilang lengkap. Karena bahasa ini open source, banyak sekali developer yabg juga mengembangkan modul-modul lain yang bisa dimanfaatkan.

### Number 2
Golang adalah bahasa pemrograman open source yang dikembankgan oleh tim google. Tetapi mempertahankan kekuatan bahasa pemrograman modern saat mengembangkan apilikasi dan jadilah golang.

Struktur Pemrograman Golang 
1.Dekralasi package
2.import library
3.Bagian Fungsi Utama / Method main.
4.Ekspresi

-Deklarasi Package adalah suatu cara untuk mengoraganisasikan file yang ada atau bisa di bilang nama folder yang tersedia

-Import library adalah memanggil library yang ingin kita gunakan fungsinya untuk memudahkan dalam membuat suatu program

-Method main adalah eksekusi pertama kali saat program tersebut dijalankan

-Ekspresi adalah statement yang merupakan sebuah kode bagian terkecil.

### Number 3

Tipe data di golang 5 yaitu :
1.tipe data numeric : 
-uint8 = 0 ↔ 255
- uint16 = 0 ↔ 65535
- uint32 = 0 ↔ 4294967295
- uint64 = 0 ↔ 18446744074709551615
- byte = sama dengan uint8
- int8 = -128 ↔  127
-  int16 = -32768 ↔  32767
-  int32 = -2147483648 ↔  2147483647
-  int64 = -9223372036854775808 ↔ 9223372036854775807
-  int = sama dengan int32 atau int64 (tergantung nilainya)
  
2.tipe data numeric desimal : float32 dan float64. perbedaan kedua tipe data tersebut adalah lebar cakupan nilai desimal yang bisa ditampung

3.tipe data boolean : true dan false

4.tipe data string : 
- isi dari string bisa 0 sampai tak terhingga
-Ciri khas data string adalah nilainya di apit oleh tanda quote atau petik dua (")

5.tipe data nill&zero value 


contoh konversi nilai: 
package main 

import (
	"fmt"
)

func main(){
    var name  = "Arman"
    var e byte = name[0]
	var eString string =  string(e)
	
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}

cara merubah atau mengkorversi tipe data di golang menggunakan funsi seperti:
-strconv.Atoi() untuk mengkoversi tipe data string ke interger
-strconv.Itoa() untuk mengkoversi tipe data interger ke string
-strconv.ParseInt() untuk mengkoversi string ke tipe numerik dengan ketentuan basis yang digunakan 10 dan lebar datanya mengikuti tipe int64
-strconv.FormatInt () untuk mengkoversi data numerik int64 ke string dengan basis bisa ditentukan sendiri
-strconv.ParseFloat() digunakan untuk mengkoversi string ke numerik desimal dengan lebar yang bisa ditenrukan sendiri
-strconv.FormatFloat()digunakan untuk mengkonversi data bertipe float64 ke string dengan format eksponen dengan lebar digit desimal dan lebar tipe data yang bisa ditentukan sendiri
-strconv.ParseBool()  digunakan untuk mengkonversi string ke bool
-strconv.FormatBool() digunakan untuk konversi bool ke string

### Number 4

Pseudocode :
{
    variable i = 1;
    for i <= 20;
    jika i modulus 2 = 0{
        maka akan menampilkan "Genap"
    } jika i modulus 2 bukan = 0 {
        maka akan menampilkan "Ganjil"
    }jika tidak {
        menampilkan "something go wrong"
    }
}

### Number 5

- Array adalah kumpulan data yang bertipe data sama, yang disimpan dalam sebuah variabel. Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatanya, menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan. Jika integer maka tiap element zero value0nya adalah 0, jika bool maka false, dan seterusnya. Setiap elemen array memiliki indeks berupa angka yang merepresentassikan posisi urutan elemen tersebut.Daya tampung array tidak bisa bertambah setelah array dibuat
- Contoh :
Package main

import (
    "fmt"
)

func main(){
    var names[2]string
    names[0] = "Arman"
    names[1] = "Zulfikri"
    names[2] = "Fardiansyah"

    fmt.Println(names[0])
    fmt.Println(names[1])
    fmt.Println(names[2])
}

-Slice adalah potongan dari data Array, Slice mirip dengan Array, yang membedakanya adalah ukuran slice bisa berubah, Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di-Array

Contoh : 
package main

import(
    "fmt"
)

func main(){
    var months = [...]string{
        "januari","februari","maret","april","mei",
        "juni","juli,"agustus","oktober","november","desember"
    }
    var slice1 = months [4:6]
    fmt.Println(slice1)
}

-Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. untuk setiap data (atau value) yang disimpan,disiapkan juga key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.

Perbedaan map dengan slice dapat dilihat dari indeks yang digunakan pengaksesan bisa ditentukan sendiri tipe-nya( indeks tersebut adalah key)

Contoh :
package main

func main() {
    person := map[string]string{
        "name":"Arman",
        "asal" : "Banyuwangi"
    }

    fmt.Println(person)
    fmt.println(person["name"])
}