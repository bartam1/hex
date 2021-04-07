package command

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
)

type MakeUrlHashModel interface {
	MakeUrlHash(ctx context.Context, mu domain.MakeUrlHash) (u domain.UrlHash, err error)
}

type MakeUrlHash struct {
	writeModel MakeUrlHashModel
}

func NewMakeUrlHash(model MakeUrlHashModel) MakeUrlHash {
	return MakeUrlHash{writeModel: model}
}

func (h MakeUrlHash) Do(ctx context.Context, u domain.MakeUrlHash) (user domain.UrlHash, err error) {
	//TODO: generate hash inside core domain
	return h.writeModel.MakeUrlHash(ctx, u)
}
