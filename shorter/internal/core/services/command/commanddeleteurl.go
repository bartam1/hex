package command

import (
	"context"
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

func (h DeleteUrl) Do(ctx context.Context, hash string) error {
	return h.writeModel.DeleteUrl(ctx, hash)
}
