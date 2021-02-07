package seeders

import (
	"CasePoint3/models"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

//SeedPerson func
func SeedPerson(db *gorm.DB) {
	var SeedPersonArray = [...][10]string{
		{"1", "1234567823", "Arman Zulfikri", "Arman", "Zulfikri", "1997-06-20", "Banyuwangi", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WIB"},
		{"2", "678126386712", "Nandia Putri", "Nandia", "Putri", "1998-03-12", "Yogya", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "F", "WIB"},
		{"3", "301737821", "Reza Rahardian", "Reza", "Rahardian", "1999-12-12", "Jakarta", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WIB"},
		{"4", "12938712", "Ezza Nanda", "Ezza", "Nanda", "1997-01-20", "Lampung", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WIB"},
		{"5", "239769126321", "Syabiq Abdillah", "Syabiq", "Abdillah", "1997-09-12", "Samarinda", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WITA"},
		{"6", "213128", "Yudhi Adhi", "Yudhi", "Adhi", "1998-02-13", "Temanggung", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WIB"},
		{"7", "812738921", "Abdul Aziz", "Abdul", "Aziz", "1995-02-18", "Lampung", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WIB"},
		{"8", "1293798021", "Winnie Connie", "Winnie", "Connie", "1994-01-12", "Bekasi", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "F", "WIB"},
		{"9", "12637812", "Eddi Wicaksono", "Eddi", "Wicaksono", "1996-08-10", "Surabaya", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WIB"},
		{"10", "8216389712", "Ilham Udin", "Ilham", "Udin", "1998-04-12", "Blitar", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "M", "WIB"},
		{"11", "382173912", "Dhea Romadhona ", "Dhea", "Romadhona", "1999-02-27", "Ponorogo", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "F", "WIB"},
		{"12", "89427149", "Aprilia Sari", "Aprilia", "Sari", "1999-03-10", "Kebumen", "https://cdn.pixabay.com/photo/2017/06/13/12/53/profile-2398782_960_720.png", "F", "WIB"},
	}

	var person models.Persons
	for _, v := range SeedPersonArray {
		subDistrict, _ := strconv.ParseUint(v[0], 10, 32)
		person.SubDistrictID = uint(subDistrict)
		person.Nip = v[1]
		person.FullName = v[2]
		person.FirstName = v[3]
		person.LastName = v[4]
		person.BirthDate = v[5]
		person.BirthPlace = v[6]
		person.Photo = v[7]
		person.Gender = v[8]
		person.ZonaLocation = v[9]
		person.ID = 0
		db.Create(&person)

	}
	fmt.Println("Seeder Persons created Sucessfully")
}
