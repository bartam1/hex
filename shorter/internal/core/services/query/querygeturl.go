package query

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
)

type UrlReadModel interface {
	GetUrl(ctx context.Context, h string) (u domain.UrlHash, err error)
}

type Url struct {
	readModel UrlReadModel
}

func NewUrl(model UrlReadModel) Url {
	return Url{readModel: model}
}

func (h Url) Do(ctx context.Context, hash string) (u domain.UrlHash, err error) {
	return h.readModel.GetUrl(ctx, hash)
}
