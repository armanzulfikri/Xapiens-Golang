package buffer

import "testing"

func BenchmarkPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//dataPostFormData()
	}
}
