package carbon

import "testing"

func BenchmarkCarbon_Closest(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Closest(Parse("2020-08-05"), Parse("2022-08-05"))
	}
}

func BenchmarkCarbon_Farthest(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Farthest(Parse("2020-08-05"), Parse("2022-08-05"))
	}
}
