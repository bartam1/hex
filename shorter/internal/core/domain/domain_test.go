package domain_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
)

func makeRandomString(max int) string {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	n := rng.Intn(max)
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
	}
	return string(runes)
}

func TestHashRnd(t *testing.T) {
	for i := 0; i < 1000; i++ {
		teststr1 := makeRandomString(i + 5)
		teststr2 := makeRandomString(i + 5)
		if domain.Hash(teststr1) == domain.Hash(teststr2) {
			t.Errorf("teststr1: %q, teststr2: %q", teststr1, teststr2)
		}
	}
}

func TestHashLen(t *testing.T) {
	for i := 0; i < 1000; i++ {
		teststr1 := makeRandomString(i + 5)
		l := len(domain.GenerateUrlHash(teststr1).Hash)
		if l != 8 {
			t.Errorf("len: %d, teststr1: %q", l, teststr1)
		}
	}
}
