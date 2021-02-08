package controller

import (
	"fmt"

	"gorm.io/gorm"
)

//SumGender func
func SumGender(db *gorm.DB) {
	type NResult struct {
		N int64
	}

	var n NResult

	db.Table("persons").Select("sum(gender) as g").Scan(&n)

	fmt.Println(n.N)
}
