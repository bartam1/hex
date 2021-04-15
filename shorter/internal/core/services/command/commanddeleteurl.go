package command

import (
	"context"

	"github.com/bartam1/mobilfox/shorter/pkg/logs/extlog"
)

type DeleteUrlModel interface {
	DeleteUrl(ctx context.Context, hash string) error
}

type DeleteUrl struct {
	writeModel DeleteUrlModel
}

func NewDeleteUrl(model DeleteUrlModel) DeleteUrl {
	return DeleteUrl{writeModel: model}
}

func (h DeleteUrl) Do(ctx context.Context, hash string) (err error) {
	defer func() {
		extlog.LogCommandExecution("DeleteUrl", hash, err)
	}()
	return h.writeModel.DeleteUrl(ctx, hash)
}
