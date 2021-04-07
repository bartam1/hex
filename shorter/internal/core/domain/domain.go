package domain

import (
	"hash/fnv"
	"strconv"
)

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func GenerateUrlHash(url string) UrlHash {
	hstring := strconv.FormatInt(int64(hash(url)), 16)
	return UrlHash{Url: url, Hash: hstring[1:9]}
}
