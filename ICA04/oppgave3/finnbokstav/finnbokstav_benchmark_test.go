package oppg3

import (
	
	"testing"
	"math/rand"
	"time"
)

var liten string = "text1.txt"
var middels string = "Middels.txt"
var stor string = "pg100.txt"

func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

func BenchmarkFinnLiten(b *testing.B) {
	Finn(liten, b)

}

func BenchmarkFinnMiddels(b *testing.B) {
	Finn(middels, b)

}

func BenchmarkFinnStor(b *testing.B) {
	Finn(stor, b)

}

func Finn(i string, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		b.StartTimer()
		FinnAlletest(i)

	}
}
