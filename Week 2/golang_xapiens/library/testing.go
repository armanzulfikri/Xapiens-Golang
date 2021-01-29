package library

import (
	"testing"
)
var (
	volumeSeharusnya float64 = 125
)

func TestingHitungVolume(t testing.T){
	t.Logf("volume :  %.2f", KubusVolume(5))
	
	if KubusVolume(5) != volumeSeharusnya {
	t.Errorf("Salah, seharusnya : %.2f", volumeSeharusnya)
	}
}


func KubusVolume(angka float64)(volume float64){
	volume = angka * angka
	return
}