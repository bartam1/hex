package command

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
)

type MakeUrlHashModel interface {
	MakeUrlHash(ctx context.Context, mu domain.UrlHash) (u domain.UrlHash, err error)
}

type MakeUrlHash struct {
	writeModel MakeUrlHashModel
}

func NewMakeUrlHash(model MakeUrlHashModel) MakeUrlHash {
	return MakeUrlHash{writeModel: model}
}

func (h MakeUrlHash) Do(ctx context.Context, u domain.MakeUrlHash) (urlh domain.UrlHash, err error) {
	uh := domain.GenerateUrlHash(u.Url)
	return h.writeModel.MakeUrlHash(ctx, uh)
}
