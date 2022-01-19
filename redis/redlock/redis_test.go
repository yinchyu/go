package compent

import (
	"fmt"
	"testing"
)

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
func TestClient(t *testing.T) {
	cmd := client.Do("Set", 5, "北京")
	fmt.Println(cmd.String())
	fmt.Println(client.Get("5"))
}
