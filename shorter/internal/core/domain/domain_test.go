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
	puff := make([]domain.UrlHash, 1000)
	var urlhash domain.UrlHash
	for i := 0; i < 1000; i++ {
		teststr1 := makeRandomString(i + 5)
		urlhash = domain.GenerateUrlHash(teststr1)
		l := len(urlhash.Hash)
		if l != 8 {
			t.Errorf("len: %d, teststr1: %q too long hash string!", l, teststr1)
		}
		puff = append(puff, urlhash)
	}
	for i := 0; i < 1000; i++ {
		teststr1 := makeRandomString(i + 5)
		urlhash = domain.GenerateUrlHash(teststr1)
		l := len(urlhash.Hash)
		if l != 8 {
			t.Errorf("len: %d, teststr1: %q too long hash string!", l, teststr1)
		}
		for _, h := range puff {
			if urlhash.Hash == h.Hash && urlhash.Url != h.Url {
				t.Errorf("teststr1: %q, teststr2: %q - Both url have same hash!", urlhash.Url, h.Url)
			}
		}
	}
}
