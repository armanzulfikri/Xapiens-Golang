package library

import "testing"

var (volumeSeharusnya float64 = 15)


//TestHitungVolume Example


func TestHitungVolume(t *testing.T){

	t.Logf("volume :  %.2f", KubusVolume(5))


	if KubusVolume(5) != volumeSeharusnya {
		t.Errorf("Salah, seharusnya : %.2f", volumeSeharusnya)
	}



}