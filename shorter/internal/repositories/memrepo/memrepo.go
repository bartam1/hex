package memrepo

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
)

type memDB struct {
	db map[string]string
}

func (d memDB) GetUrlsWidthHash(ctx context.Context) (us []domain.UrlHash, err error) {
	for key, element := range d.db {
		us = append(us, domain.UrlHash{Url: element, Hash: key})
	}
	return us, nil
}
func (d memDB) GetUrl(ctx context.Context, url string) (u domain.UrlHash, err error) {
	return domain.UrlHash{Url: d.db[url], Hash: url}, nil
}
func (d memDB) MakeUrlHash(ctx context.Context, mu domain.UrlHash) (u domain.UrlHash, err error) {
	d.db[mu.Hash] = mu.Url
	return mu, nil
}
func (d memDB) DeleteUrl(ctx context.Context, hash string) error {
	return nil
}
func New() (memDB, error) {
	return memDB{db: map[string]string{}}, nil
}
