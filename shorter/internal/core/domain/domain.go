package domain

import (
	"crypto/sha1"
	"io"
)

func GenerateUrlHash(url string) UrlHash {
	h := sha1.New()
	io.WriteString(h, url)
	hash := h.Sum(nil)
	return UrlHash{Url: url, Hash: string(hash[:8])}
}
