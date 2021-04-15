package command

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
	"github.com/bartam1/mobilfox/shorter/pkg/logs/extlog"
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
	defer func() {
		extlog.LogCommandExecution("MakeUrlHash", u.Url, err)
	}()
	uh := domain.GenerateUrlHash(u.Url)
	return h.writeModel.MakeUrlHash(ctx, uh)
}
