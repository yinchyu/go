package compent

import "testing"

func BenchmarkAuthCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AuthCode(6)
	}
}
func BenchmarkAuthCode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AuthCode2()
	}
}
