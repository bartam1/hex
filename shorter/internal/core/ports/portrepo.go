package port

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
)

type RepoIf interface {
	GetUrlsWidthHash(ctx context.Context) (us []domain.UrlHash, err error)
	GetUrl(ctx context.Context, url string) (u domain.UrlHash, err error)
	MakeUrlHash(ctx context.Context, mu domain.UrlHash) (u domain.UrlHash, err error)
	DeleteUrl(ctx context.Context, hash string) error
}
