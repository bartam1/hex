package memrepo

import (
	"context"

	"errors"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
	"github.com/bartam1/mobilfox/shorter/pkg/errors/exterror"
	"github.com/sirupsen/logrus"
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
func (d memDB) GetUrl(ctx context.Context, hash string) (u domain.UrlHash, err error) {
	if d.db[hash] == "" {
		return domain.UrlHash{}, exterror.NewRepoSlug(errors.New("No url hash"), "There is no url with that hash!", func() { logrus.Error(errors.New("No url with that hash")) })
	}
	return domain.UrlHash{Url: d.db[hash], Hash: hash}, nil
}

func (d memDB) MakeUrlHash(ctx context.Context, mu domain.UrlHash) (u domain.UrlHash, err error) {
	d.db[mu.Hash] = mu.Url
	return mu, nil
}
func (d memDB) DeleteUrl(ctx context.Context, hash string) error {
	if d.db[hash] == "" {
		return exterror.NewRepoSlug(errors.New("No url hash"), "There is no url with that hash!", func() { logrus.Error(errors.New("No url hash")) })
	}
	delete(d.db, hash)
	return nil
}
func New() (memDB, error) {
	return memDB{db: map[string]string{}}, nil
}
