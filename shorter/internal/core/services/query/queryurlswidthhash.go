package query

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
)

type UrlsWidthHashReadModel interface {
	GetUrlsWidthHash(ctx context.Context) (us []domain.UrlHash, err error)
}

type UrlsWidthHash struct {
	readModel UrlsWidthHashReadModel
}

func NewUrlsWidthHash(model UrlsWidthHashReadModel) UrlsWidthHash {
	return UrlsWidthHash{readModel: model}
}

func (h UrlsWidthHash) Do(ctx context.Context) (us []domain.UrlHash, err error) {
	return h.readModel.GetUrlsWidthHash(ctx)
}
