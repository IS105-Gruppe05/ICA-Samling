package oppg3

import (
	//"fmt"
	"testing"
	//"reflect"
	"time"
	"math/rand"

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


func BenchmarkFinnBufferLiten(b *testing.B) {
	FinnBuffer(liten, b)
	// Output: MOOOO!

}
func BenchmarkFinnBufferMiddels(b *testing.B) {
	FinnBuffer(middels, b)
	// Output: MOOOO!

}

func BenchmarkFinnBufferStor(b *testing.B) {
	FinnBuffer(stor, b)
	// Output: MOOOO!

}

func FinnBuffer(i string, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		//values := perm(i)
		b.StartTimer()
		ReadBufferTest(i)

	}
}
